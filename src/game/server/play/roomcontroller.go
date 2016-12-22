package play

import (
	. "types"
)

func Handle_LoginRoomReq(objectId IdString, opCode MsgType, req *LoginRoomReq) interface{} {
	ack := &LoginRoomAck{}

	player, ok := OnlineM.GetOnlinePlayer(objectId)
	if !ok {
		ack.Common = getCommonAck(ERR_PLAYER_NOT_FOUND)
		return ack
	}

	room, ok := RoomM.FindRoom(IdString(req.RoomId))
	if !ok {
		room = RoomM.CreateRoom(player.UserId, player.Name)
		room.BeginFrameSync()
	}

	sa := &ScrActive{
		Id:        player.UserId,
		BelongsTo: room.BrdCastGroup.Id,
		Name:      player.Name,

		Type:    screenObjTypePlayer,
		SubType: player.SubType,
		Skin:    player.Skin,
		Hp:      player.Hp,
		FullHp:  player.FullHp,
	}
	sa.PosX = int16(req.PosX)
	sa.PosY = int16(req.PosY)

	room.AddOrReplaceActive(sa)
	room.AddBcgMember(player.UserId)

	bcgSync := &BrdCastGroupManageReq{
		GroupId:   string(room.BrdCastGroup.Id),
		MemberIds: []string{string(player.UserId)},
	}
	AsyncSender.SendServerNotify(MT_BrdCastAddMemberReq, bcgSync)

	ack.Screen = room.Screen.ToClient()
	ack.Common = getCommonAck(OK)

	return ack
}

func Handle_LeaveRoomReq(objectId IdString, opCode MsgType, req *LeaveRoomReq) interface{} {
	ack := &LeaveRoomAck{}

	player, ok := OnlineM.GetOnlinePlayer(objectId)
	if !ok {
		ack.Common = getCommonAck(ERR_PLAYER_NOT_FOUND)
		return ack
	}

	room, ok := RoomM.FindRoom(IdString(req.RoomId))
	if !ok {
		ack.Common = getCommonAck(ERR_ROOM_NOT_FOUND)
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

	ack.Common = getCommonAck(OK)
	return ack
}
