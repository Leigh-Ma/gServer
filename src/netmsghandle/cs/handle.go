package cs
//Auto generated, do not modify unless you know clearly what you are doing.

import (
	pb "github.com/golang/protobuf/proto"
	. "types"
	"library/logger"
	"library/structenh"
    . "game/server/play"
)

func On_LoginReq(objectId IdString, opCode MsgType, payLoad []byte) interface{} {
	req := &LoginReq{}
	pb.Unmarshal(payLoad, req)
	logger.Payload("rx-r: %4d: payload %s", opCode, structenh.StringifyStruct(req))
	return Handle_LoginReq(objectId, opCode, req)
}

func On_LoginAck(objectId IdString, opCode MsgType, payLoad []byte) interface{} {
	req := &LoginAck{}
	pb.Unmarshal(payLoad, req)
	logger.Payload("rx-r: %4d: payload %s", opCode, structenh.StringifyStruct(req))
	return Handle_LoginAck(objectId, opCode, req)
}

func On_LogoutReq(objectId IdString, opCode MsgType, payLoad []byte) interface{} {
	req := &LogoutReq{}
	pb.Unmarshal(payLoad, req)
	logger.Payload("rx-r: %4d: payload %s", opCode, structenh.StringifyStruct(req))
	return Handle_LogoutReq(objectId, opCode, req)
}

func On_LogoutAck(objectId IdString, opCode MsgType, payLoad []byte) interface{} {
	req := &LogoutAck{}
	pb.Unmarshal(payLoad, req)
	logger.Payload("rx-r: %4d: payload %s", opCode, structenh.StringifyStruct(req))
	return Handle_LogoutAck(objectId, opCode, req)
}

func On_SetPlayerNameReq(objectId IdString, opCode MsgType, payLoad []byte) interface{} {
	req := &SetPlayerNameReq{}
	pb.Unmarshal(payLoad, req)
	logger.Payload("rx-r: %4d: payload %s", opCode, structenh.StringifyStruct(req))
	return Handle_SetPlayerNameReq(objectId, opCode, req)
}

func On_SetPlayerNameAck(objectId IdString, opCode MsgType, payLoad []byte) interface{} {
	req := &SetPlayerNameAck{}
	pb.Unmarshal(payLoad, req)
	logger.Payload("rx-r: %4d: payload %s", opCode, structenh.StringifyStruct(req))
	return Handle_SetPlayerNameAck(objectId, opCode, req)
}

func On_ChoseSideReq(objectId IdString, opCode MsgType, payLoad []byte) interface{} {
	req := &ChoseSideReq{}
	pb.Unmarshal(payLoad, req)
	logger.Payload("rx-r: %4d: payload %s", opCode, structenh.StringifyStruct(req))
	return Handle_ChoseSideReq(objectId, opCode, req)
}

func On_ChoseSideAck(objectId IdString, opCode MsgType, payLoad []byte) interface{} {
	req := &ChoseSideAck{}
	pb.Unmarshal(payLoad, req)
	logger.Payload("rx-r: %4d: payload %s", opCode, structenh.StringifyStruct(req))
	return Handle_ChoseSideAck(objectId, opCode, req)
}

func On_StartFightReq(objectId IdString, opCode MsgType, payLoad []byte) interface{} {
	req := &StartFightReq{}
	pb.Unmarshal(payLoad, req)
	logger.Payload("rx-r: %4d: payload %s", opCode, structenh.StringifyStruct(req))
	return Handle_StartFightReq(objectId, opCode, req)
}

func On_StartFightAck(objectId IdString, opCode MsgType, payLoad []byte) interface{} {
	req := &StartFightAck{}
	pb.Unmarshal(payLoad, req)
	logger.Payload("rx-r: %4d: payload %s", opCode, structenh.StringifyStruct(req))
	return Handle_StartFightAck(objectId, opCode, req)
}

func On_LeaveRoomReq(objectId IdString, opCode MsgType, payLoad []byte) interface{} {
	req := &LeaveRoomReq{}
	pb.Unmarshal(payLoad, req)
	logger.Payload("rx-r: %4d: payload %s", opCode, structenh.StringifyStruct(req))
	return Handle_LeaveRoomReq(objectId, opCode, req)
}

func On_LeaveRoomAck(objectId IdString, opCode MsgType, payLoad []byte) interface{} {
	req := &LeaveRoomAck{}
	pb.Unmarshal(payLoad, req)
	logger.Payload("rx-r: %4d: payload %s", opCode, structenh.StringifyStruct(req))
	return Handle_LeaveRoomAck(objectId, opCode, req)
}

func On_MoveActionReq(objectId IdString, opCode MsgType, payLoad []byte) interface{} {
	req := &MoveActionReq{}
	pb.Unmarshal(payLoad, req)
	logger.Payload("rx-r: %4d: payload %s", opCode, structenh.StringifyStruct(req))
	return Handle_MoveActionReq(objectId, opCode, req)
}

func On_MoveActionAck(objectId IdString, opCode MsgType, payLoad []byte) interface{} {
	req := &MoveActionAck{}
	pb.Unmarshal(payLoad, req)
	logger.Payload("rx-r: %4d: payload %s", opCode, structenh.StringifyStruct(req))
	return Handle_MoveActionAck(objectId, opCode, req)
}

func On_ChoseHeroReq(objectId IdString, opCode MsgType, payLoad []byte) interface{} {
	req := &ChoseHeroReq{}
	pb.Unmarshal(payLoad, req)
	logger.Payload("rx-r: %4d: payload %s", opCode, structenh.StringifyStruct(req))
	return Handle_ChoseHeroReq(objectId, opCode, req)
}

func On_ChoseHeroAck(objectId IdString, opCode MsgType, payLoad []byte) interface{} {
	req := &ChoseHeroAck{}
	pb.Unmarshal(payLoad, req)
	logger.Payload("rx-r: %4d: payload %s", opCode, structenh.StringifyStruct(req))
	return Handle_ChoseHeroAck(objectId, opCode, req)
}

func On_SearchRoomReq(objectId IdString, opCode MsgType, payLoad []byte) interface{} {
	req := &SearchRoomReq{}
	pb.Unmarshal(payLoad, req)
	logger.Payload("rx-r: %4d: payload %s", opCode, structenh.StringifyStruct(req))
	return Handle_SearchRoomReq(objectId, opCode, req)
}

func On_SearchRoomAck(objectId IdString, opCode MsgType, payLoad []byte) interface{} {
	req := &SearchRoomAck{}
	pb.Unmarshal(payLoad, req)
	logger.Payload("rx-r: %4d: payload %s", opCode, structenh.StringifyStruct(req))
	return Handle_SearchRoomAck(objectId, opCode, req)
}

func On_ScreenChangeNotify(objectId IdString, opCode MsgType, payLoad []byte) interface{} {
	req := &ScreenChangeNotify{}
	pb.Unmarshal(payLoad, req)
	logger.Payload("rx-r: %4d: payload %s", opCode, structenh.StringifyStruct(req))
	return Handle_ScreenChangeNotify(objectId, opCode, req)
}

func On_StartFightNotify(objectId IdString, opCode MsgType, payLoad []byte) interface{} {
	req := &StartFightNotify{}
	pb.Unmarshal(payLoad, req)
	logger.Payload("rx-r: %4d: payload %s", opCode, structenh.StringifyStruct(req))
	return Handle_StartFightNotify(objectId, opCode, req)
}

func On_ChoseSideNotify(objectId IdString, opCode MsgType, payLoad []byte) interface{} {
	req := &ChoseSideNotify{}
	pb.Unmarshal(payLoad, req)
	logger.Payload("rx-r: %4d: payload %s", opCode, structenh.StringifyStruct(req))
	return Handle_ChoseSideNotify(objectId, opCode, req)
}
