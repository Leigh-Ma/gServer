package play

import (
	"library/database"
	. "types"
)

func Handle_LoginReq(objectId IdString, opCode MsgType, req *LoginReq) interface{} {
	ack := &LoginAck{}
	//FORCE reload player from database, player data maybe changed in other server
	player, ok := AllPlayerM.LoadOneFromDb(database.MongoProxy, objectId)

	if !ok {
		player = &Player{
			UserId: objectId,
			Name:   "", //new user
			UUID:   req.Uuid,
		}
		player.Hero = *NewHero()
		AllPlayerM.AddPlayer(player)
		database.DbUpsert(player)
	}

	hero := player.Hero

	ack.Common = getCommonAck(OK)
	ack.HeroType = hero.HeroType
	ack.HeroName = hero.HeroName
	ack.HeroSkills = hero.Skills
	ack.HeroSkin = hero.Skin
	ack.UserName = player.Name
	ack.UserId = string(player.UserId)

	return ack
}

func Handle_LogoutReq(objectId IdString, opCode MsgType, req *LogoutReq) interface{} {
	player, ok := AllPlayerM.GetPlayer(objectId)
	if !ok {
		return nil
	}

	if player.room != nil {
		player.room.DelBcgMember(player.UserId)
		bcgSync := &BrdCastGroupManageReq{
			GroupId:   string(player.room.BrdCastGroup.Id),
			MemberIds: []string{string(player.UserId)},
		}
		AsyncSender.SendServerNotify(MT_BrdCastDelMemberReq, bcgSync)
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

	player.Name = req.Name
	if player.room != nil {
		sa := player.room.GetExistActive(player.UserId)
		if sa != nil {
			sa.Name = player.Name
			sa.SetDetailChanged()
			player.room.AddOrReplaceActive(sa)
		}
	}
	ack.Name = player.Name
	ack.Common = getCommonAck(OK)

	database.DbUpsert(player)
	return ack
}
