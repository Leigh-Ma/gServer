package play

import (
	. "types"
)

func Handle_LoginReq(objectId IdString, opCode MsgType, req *LoginReq) interface{} {

	ack := &LoginAck{Common: getCommonAck(OK)}

	return ack
}

func Handle_LogoutReq(objectId IdString, opCode MsgType, req *LogoutReq) interface{} {
	return nil
}
