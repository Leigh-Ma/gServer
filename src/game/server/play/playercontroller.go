package play

import (
	"library/database"
	. "types"
)

func Handle_LoginReq(objectId IdString, opCode MsgType, req *LoginReq) interface{} {
	ack := &LoginAck{}
	player, ok := OnlineM.GetOnePlayer(objectId)
	if !ok {
		player, ok = AllPlayerM.GetOnePlayer(objectId)
	}
	if !ok {
		player = &Player{
			UserId: objectId,
			Name:   RandomPlayerName(),
			UUID:   req.Uuid,
		}
		player.Hero = *NewHero()
	}

	hero := player.Hero
	//query or create from db
	OnlineM.AddOnePlayer(player)
	ack.Common = getCommonAck(OK)
	ack.HeroName = hero.HeroName
	ack.HeroSkills = hero.Skills
	ack.HeroSkin = hero.Skin
	ack.UserName = player.Name
	ack.UserId = string(player.UserId)

	database.DbUpsert(player)
	return ack
}

func Handle_LogoutReq(objectId IdString, opCode MsgType, req *LogoutReq) interface{} {
	player, ok := OnlineM.GetOnePlayer(objectId)
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

	OnlineM.DelOnePlayer(objectId)
	return nil
}

func Handle_SetPlayerNameReq(objectId IdString, opCode MsgType, req *SetPlayerNameReq) interface{} {
	ack := &SetPlayerNameAck{}
	player, ok := OnlineM.GetOnePlayer(objectId)
	if ok {
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
	ack.Common = getCommonAck(OK)

	database.DbUpsert(player)
	return ack
}
