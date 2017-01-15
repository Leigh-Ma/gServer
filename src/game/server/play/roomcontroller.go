package play

import (
	"library/config"
	"library/database"
	"math/rand"
	. "types"
)

func Handle_SearchRoomReq(objectId IdString, opCode MsgType, req *SearchRoomReq) interface{} {
	ack := &SearchRoomAck{}

	//FORCE reload player from database, player data maybe changed in other server
	player, ok := AllPlayerM.LoadOneFromDb(database.MongoProxy, objectId)
	if !ok {
		ack.Common = getCommonAck(ERR_PLAYER_NOT_FOUND)
		return ack
	}
	room := RoomM.ChoseByRoomId(req.RoomId)
	if room == nil {
		room = RoomM.CreateRoom(player.UserName)
		room.BeginFrameSync() //will create broadcast group on gateway
	}

	ack.RoomId = string(room.Id)
	ack.RoomName = room.Name
	ack.Red, ack.Blue = room.SideRoles()

	//fake data for debug
	ack.Red = append(ack.Red, NewHero().debugRoleInfo())
	ack.Red = append(ack.Red, NewHero().debugRoleInfo())
	ack.Blue = append(ack.Blue, NewHero().debugRoleInfo())
	ack.Blue = append(ack.Blue, NewHero().debugRoleInfo())
	ack.Blue = append(ack.Blue, NewHero().debugRoleInfo())

	ack.Common = getCommonAck(OK)

	return ack
}

func Handle_ChoseSideReq(objectId IdString, opCode MsgType, req *ChoseSideReq) interface{} {
	ack := &ChoseSideAck{}
	isNewRoom := false
	//FORCE reload player from database, player data maybe changed in other server
	player, ok := AllPlayerM.LoadOneFromDb(database.MongoProxy, objectId)
	if !ok {
		ack.Common = getCommonAck(ERR_PLAYER_NOT_FOUND)
		return ack
	}

	room := RoomM.ChoseByRoomId(req.RoomId)

	sa := room.GetExistActive(player.UserId)
	if sa == nil {
		sa = &ScrActive{
			Id:   player.UserId,
			Type: screenObjTypePlayer,
			Side: config.GetRoomSideNo(req.Side),
		}
	}

	sa.PosX = int16(rand.Int31n(scrWidth))
	sa.PosY = int16(rand.Int31n(scrHeight))

	player.FillActiveDetail(&sa.ActDetail)
	sa.SetDetailChanged()

	room.NoNotifyAddActive(sa)
	room.AddBcgMember(player.UserId)

	player.room = room

	if isNewRoom {
		room.BeginFrameSync()
	} else {
		bcgSync := &BrdCastGroupManageReq{
			GroupId:   string(room.BrdCastGroup.Id),
			MemberIds: []string{string(player.UserId)},
		}
		AsyncSender.SendServerNotify(MT_BrdCastAddMemberReq, bcgSync)
	}

	notify := &ChoseSideNotify{Side: req.Side, Room: room.RoomInfo(), Role: sa.RoleInfo()}
	AsyncSender.InstantSendBroadCastNotify(MT_ChoseSideNotify, room.Id, notify)

	ack.Screen = room.ScreenInfo()
	ack.Room = room.RoomInfo()
	ack.Common = getCommonAck(OK)

	database.DbUpsert(player)

	return ack
}

func Handle_LeaveRoomReq(objectId IdString, opCode MsgType, req *LeaveRoomReq) interface{} {
	ack := &LeaveRoomAck{}

	player, ok := AllPlayerM.GetPlayer(objectId)
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
	player, ok := AllPlayerM.GetPlayer(objectId)
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

func Handle_StartFightReq(objectId IdString, opCode MsgType, req *StartFightReq) interface{} {
	ack := &StartFightAck{}

	player, ok := AllPlayerM.GetPlayer(objectId)
	if !ok {
		ack.Common = getCommonAck(ERR_PLAYER_NOT_FOUND)
		return ack
	}

	room, ok := RoomM.FindRoom(IdString(req.RoomId))
	if !ok {
		ack.Common = getCommonAck(ERR_ROOM_NOT_FOUND)
		return ack
	}

	sa := room.GetExistActive(player.UserId)
	if sa == nil {
		ack.Common = getCommonAck(ERR_SHOULD_LOGIN_ROOM)
		return ack
	}

	room.FightBegin()

	ack.Common = getCommonAck(OK)
	return ack
}
