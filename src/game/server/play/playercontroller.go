package play

import (
	"library/config"
	"library/database"
	"library/logger"
	. "types"
)

func Handle_LoginReq(objectId IdString, opCode MsgType, req *LoginReq) interface{} {
	ack := &LoginAck{}
	//FORCE reload player from database, player data maybe changed in other server
	player, ok := AllPlayerM.LoadOneFromDb(database.MongoProxy, objectId)

	if !ok {
		player = &Player{
			UserId: objectId,
			UUID:   req.Uuid,
		}
		player.Hero = *NewHero()
		AllPlayerM.AddPlayer(player)
		database.DbUpsert(player)
	}

	ack.Common = getCommonAck(OK)
	ack.Hero = player.RoleInfo()

	return ack
}

func Handle_LogoutReq(objectId IdString, opCode MsgType, req *LogoutReq) interface{} {
	player, ok := AllPlayerM.GetPlayer(objectId)
	if !ok {
		return nil
	}
	logger.Info("player %s logout, room %v", objectId, player.room)

	if player.room != nil {
		player.room.DelBcgMember(player.UserId)
		bcgSync := &BrdCastGroupManageReq{
			GroupId:   string(player.room.BrdCastGroup.Id),
			MemberIds: []string{string(player.UserId)},
		}
		AsyncSender.InstantSendServerNotify(MT_BrdCastDelMemberReq, bcgSync)
		player.room = nil
	}

	return nil
}

func Handle_SetPlayerNameReq(objectId IdString, opCode MsgType, req *SetPlayerNameReq) interface{} {
	ack := &SetPlayerNameAck{}
	player, ok := AllPlayerM.GetPlayer(objectId)
	if !ok {
		ack.Common = getCommonAck(ERR_PLAYER_NOT_FOUND)
		return ack
	}

	player.UserName = req.UserName
	if player.room != nil {
		sa := player.room.GetExistActive(player.UserId)
		if sa != nil {
			sa.Name = player.UserName
			sa.SetDetailChanged()
			player.room.AddOrReplaceActive(sa)
		}
	}
	ack.UserName = player.UserName
	ack.Common = getCommonAck(OK)

	database.DbUpsert(player)
	return ack
}

func Handle_ChoseHeroReq(objectId IdString, opCode MsgType, req *ChoseHeroReq) interface{} {
	ack := &ChoseHeroAck{}
	player, ok := AllPlayerM.GetPlayer(objectId)
	if !ok {
		ack.Common = getCommonAck(ERR_PLAYER_NOT_FOUND)
		return ack
	}

	for _, skill := range req.HeroSkills {
		if !config.IsValidHeroSkill(skill) {
			ack.Common = getCommonAck(ERR_SKILL_NOT_EXIST)
			return ack
		}
	}

	hero, ok := AllHeroes[req.HeroType]
	if !ok {
		ack.Common = getCommonAck(ERR_HERO_NOT_EXIST)
		return ack
	}

	//keep user name
	hero.UserName = player.UserName
	hero.Skills = make([]string, len(req.HeroSkills))
	copy(hero.Skills, req.HeroSkills)

	// change hero during fight, regulate hp
	if !req.ResetHp && player.HeroType != hero.HeroType {
		hero.Hp = hero.FullHp * player.Hp / player.FullHp
	}

	hero.Skin = req.HeroSkin

	//set player hero after all attributes are set
	player.Hero = hero

	room := player.room
	if room != nil {
		sa := room.GetExistActive(player.UserId)
		if sa != nil {
			player.FillActiveDetail(&sa.ActDetail)
			sa.SetDetailChanged()
			room.AddOrReplaceActive(sa)
		}
	}

	ack.Common = getCommonAck(OK)
	ack.Hero = player.RoleInfo()

	return ack
}
