package cs
//Auto generated, do not modify unless you know clearly what you are doing.

import (
	pb "github.com/golang/protobuf/proto"
	. "types"
    . "game/server/play"
)

func On_LoginReq(objectId IdString, opCode MsgType, payLoad []byte) interface{} {
	req := &LoginReq{}
	pb.Unmarshal(payLoad, req)
	return Handle_LoginReq(objectId, opCode, req)
}

func On_LoginAck(objectId IdString, opCode MsgType, payLoad []byte) interface{} {
	req := &LoginAck{}
	pb.Unmarshal(payLoad, req)
	return Handle_LoginAck(objectId, opCode, req)
}

func On_LogoutReq(objectId IdString, opCode MsgType, payLoad []byte) interface{} {
	req := &LogoutReq{}
	pb.Unmarshal(payLoad, req)
	return Handle_LogoutReq(objectId, opCode, req)
}

func On_LogoutAck(objectId IdString, opCode MsgType, payLoad []byte) interface{} {
	req := &LogoutAck{}
	pb.Unmarshal(payLoad, req)
	return Handle_LogoutAck(objectId, opCode, req)
}

func On_SetPlayerNameReq(objectId IdString, opCode MsgType, payLoad []byte) interface{} {
	req := &SetPlayerNameReq{}
	pb.Unmarshal(payLoad, req)
	return Handle_SetPlayerNameReq(objectId, opCode, req)
}

func On_SetPlayerNameAck(objectId IdString, opCode MsgType, payLoad []byte) interface{} {
	req := &SetPlayerNameAck{}
	pb.Unmarshal(payLoad, req)
	return Handle_SetPlayerNameAck(objectId, opCode, req)
}

func On_LoginRoomReq(objectId IdString, opCode MsgType, payLoad []byte) interface{} {
	req := &LoginRoomReq{}
	pb.Unmarshal(payLoad, req)
	return Handle_LoginRoomReq(objectId, opCode, req)
}

func On_LoginRoomAck(objectId IdString, opCode MsgType, payLoad []byte) interface{} {
	req := &LoginRoomAck{}
	pb.Unmarshal(payLoad, req)
	return Handle_LoginRoomAck(objectId, opCode, req)
}

func On_ChoseSideReq(objectId IdString, opCode MsgType, payLoad []byte) interface{} {
	req := &ChoseSideReq{}
	pb.Unmarshal(payLoad, req)
	return Handle_ChoseSideReq(objectId, opCode, req)
}

func On_ChoseSideAck(objectId IdString, opCode MsgType, payLoad []byte) interface{} {
	req := &ChoseSideAck{}
	pb.Unmarshal(payLoad, req)
	return Handle_ChoseSideAck(objectId, opCode, req)
}

func On_ScreenChangeNotify(objectId IdString, opCode MsgType, payLoad []byte) interface{} {
	req := &ScreenChangeNotify{}
	pb.Unmarshal(payLoad, req)
	return Handle_ScreenChangeNotify(objectId, opCode, req)
}

func On_LeaveRoomReq(objectId IdString, opCode MsgType, payLoad []byte) interface{} {
	req := &LeaveRoomReq{}
	pb.Unmarshal(payLoad, req)
	return Handle_LeaveRoomReq(objectId, opCode, req)
}

func On_LeaveRoomAck(objectId IdString, opCode MsgType, payLoad []byte) interface{} {
	req := &LeaveRoomAck{}
	pb.Unmarshal(payLoad, req)
	return Handle_LeaveRoomAck(objectId, opCode, req)
}

func On_MoveActionReq(objectId IdString, opCode MsgType, payLoad []byte) interface{} {
	req := &MoveActionReq{}
	pb.Unmarshal(payLoad, req)
	return Handle_MoveActionReq(objectId, opCode, req)
}

func On_MoveActionAck(objectId IdString, opCode MsgType, payLoad []byte) interface{} {
	req := &MoveActionAck{}
	pb.Unmarshal(payLoad, req)
	return Handle_MoveActionAck(objectId, opCode, req)
}

func On_ChoseHeroReq(objectId IdString, opCode MsgType, payLoad []byte) interface{} {
	req := &ChoseHeroReq{}
	pb.Unmarshal(payLoad, req)
	return Handle_ChoseHeroReq(objectId, opCode, req)
}

func On_ChoseHeroAck(objectId IdString, opCode MsgType, payLoad []byte) interface{} {
	req := &ChoseHeroAck{}
	pb.Unmarshal(payLoad, req)
	return Handle_ChoseHeroAck(objectId, opCode, req)
}
