package play

import (
	"library/config"
	"library/database"
	. "types"
)

func Handle_LoginRoomReq(objectId IdString, opCode MsgType, req *LoginRoomReq) interface{} {
	ack := &LoginRoomAck{}

	player, ok := OnlineM.GetOnePlayer(objectId)
	if !ok {
		ack.Common = getCommonAck(ERR_PLAYER_NOT_FOUND)
		return ack
	}

	room, ok := RoomM.FindRoom(IdString(req.RoomId))
	if !ok {
		room = RoomM.CreateRoom(player.UserId, player.Name)
		room.BeginFrameSync()
	}

	sa := room.GetExistActive(player.UserId)
	if sa == nil {
		sa = &ScrActive{
			Id:   player.UserId,
			Type: screenObjTypePlayer,
			Side: config.RoomSideMid,
		}
	}

	sa.PosX = int16(req.PosX)
	sa.PosY = int16(req.PosY)

	player.FillActiveDetail(&sa.ActDetail)
	sa.SetDetailChanged()

	room.AddOrReplaceActive(sa)
	room.AddBcgMember(player.UserId)

	player.room = room

	bcgSync := &BrdCastGroupManageReq{
		GroupId:   string(room.BrdCastGroup.Id),
		MemberIds: []string{string(player.UserId)},
	}
	AsyncSender.SendServerNotify(MT_BrdCastAddMemberReq, bcgSync)

	ack.Screen = room.Screen.ToClient()
	ack.Common = getCommonAck(OK)

	database.DbUpsert(player)

	return ack
}

func Handle_LeaveRoomReq(objectId IdString, opCode MsgType, req *LeaveRoomReq) interface{} {
	ack := &LeaveRoomAck{}

	player, ok := OnlineM.GetOnePlayer(objectId)
	if !ok {
		ack.Common = getCommonAck(ERR_PLAYER_NOT_FOUND)
		return ack
	}

	room, ok := RoomM.FindRoom(IdString(req.RoomId))
	if !ok {
		ack.Common = getCommonAck(ERR_ROOM_NOT_FOUND)
		return ack
	}

	room.DelActive(player.UserId)
	if num, ok := room.DelBcgMember(player.UserId); ok {
		bcgSync := &BrdCastGroupManageReq{
			GroupId:   string(room.BrdCastGroup.Id),
			MemberIds: []string{string(player.UserId)},
		}
		if num == 0 {
			RoomM.DestroyRoom(room)
			AsyncSender.SendServerNotify(MT_BrdCastDestroyReq, bcgSync)
		} else {
			AsyncSender.SendServerNotify(MT_BrdCastDelMemberReq, bcgSync)
		}
	}
	player.room = nil

	ack.Common = getCommonAck(OK)
	return ack
}

func Handle_MoveActionReq(objectId IdString, opCode MsgType, req *MoveActionReq) interface{} {
	ack := &MoveActionAck{}
	player, ok := OnlineM.GetOnePlayer(objectId)
	if !ok {
		ack.Common = getCommonAck(ERR_PLAYER_NOT_FOUND)
		return ack
	}

	room := player.room
	if room == nil {
		ack.Common = getCommonAck(ERR_SHOULD_LOGIN_ROOM)
		return ack
	}

	sa := room.GetExistActive(player.UserId)
	if sa == nil {
		ack.Common = getCommonAck(ERR_SHOULD_LOGIN_ROOM)
		return ack
	}

	sa.PosX = int16(req.PosX)
	sa.PosY = int16(req.PosY)
	sa.DirectX = int16(req.DirectX)
	sa.DirectY = int16(req.DirectY)
	sa.Speed = int16(req.Speed)
	sa.Hp = req.Hp
	sa.FullHp = req.FullHp

	sa.SetDetailChanged()
	room.AddOrReplaceActive(sa)
	room.AddBcgMember(player.UserId)

	return nil
}

func Handle_ChoseSideReq(objectId IdString, opCode MsgType, req *ChoseSideReq) interface{} {
	ack := &ChoseSideAck{}
	player, ok := OnlineM.GetOnePlayer(objectId)
	if !ok {
		ack.Common = getCommonAck(ERR_PLAYER_NOT_FOUND)
		return ack
	}

	room := player.room
	if room == nil {
		ack.Common = getCommonAck(ERR_SHOULD_LOGIN_ROOM)
		return ack
	}

	sa := room.GetExistActive(player.UserId)
	if sa == nil {
		ack.Common = getCommonAck(ERR_SHOULD_LOGIN_ROOM)
		return ack
	}

	sa.Side = config.GetRoomSideNo(req.Side)
	room.AddOrReplaceActive(sa)

	ack.Common = getCommonAck(OK)

	return ack
}

func Handle_ChoseHeroReq(objectId IdString, opCode MsgType, req *ChoseHeroReq) interface{} {
	ack := &ChoseHeroAck{}
	player, ok := OnlineM.GetOnePlayer(objectId)
	if !ok {
		ack.Common = getCommonAck(ERR_PLAYER_NOT_FOUND)
		return ack
	}

	if !config.IsValidHeroName(req.HeroName) {
		ack.Common = getCommonAck(ERR_HERO_NOT_EXIST)
		return ack
	}

	for _, skill := range req.HeroSkills {
		if !config.IsValidHeroSkill(skill) {
			ack.Common = getCommonAck(ERR_SKILL_NOT_EXIST)
			return ack
		}
	}

	player.HeroName = req.HeroName
	player.Skills = make([]string, len(req.HeroSkills))
	copy(player.Skills, req.HeroSkills)

	room := player.room
	if room != nil {
		sa := room.GetExistActive(player.UserId)
		if sa != nil {
			sa.Skin = req.HeroSkin
			sa.Skills = make([]string, len(req.HeroSkills))
			copy(sa.Skills, req.HeroSkills)
			sa.SubType = req.HeroName
			sa.SetDetailChanged()
			room.AddOrReplaceActive(sa)
		}
	}

	ack.Common = getCommonAck(OK)

	return ack
}
