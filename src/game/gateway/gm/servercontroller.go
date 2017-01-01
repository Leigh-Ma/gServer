package gm

//server req send to gateway, should handle by gateway, and should send no more ack
//objectId is server_id
import (
	"game/com"
	. "types"
)

func Handle_ServerLoginReq(objectId IdString, opCode MsgType, req *ServerLoginReq) interface{} {
	ack := &ServerLoginAck{Status: OK}
	com.NewBrdCastManager4Server(objectId)
	ack.Common = serverCommonAck(OK)
	return ack
}
