package gm

import (
	"game/com"
	"library/frame"
	. "types"
)

var GatewayFrame = frame.NewFrame()

func serverCommonAck(code int32) *ServerCommonAck {
	return &ServerCommonAck{ErrCode: code}
}

//equal to room, by id
var allBrdCastGrp = com.NewBrdCastGroupManage()

func RandRoomForUser(user *User) (svrId, roomId IdString) {
	return IdString("0x001"), IdString("0x002")
}
