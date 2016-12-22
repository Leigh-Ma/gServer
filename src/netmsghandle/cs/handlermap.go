package cs
//Auto generated, do not modify unless you know clearly what you are doing.
import . "types"

type NetMsgHandler func(objectId IdString, opcode MsgType, content []byte ) interface{};

type NetMsgCb struct{
	OpCode  MsgType
	RetCode MsgType
	Handler NetMsgHandler
	Desc    string
}

var NetMsgTypeHandler = map[MsgType]*NetMsgCb {
	MT_LoginReq      :&NetMsgCb{MT_LoginReq     , MT_LoginAck     , On_LoginReq     , "player login request"},
	MT_LoginAck      :&NetMsgCb{MT_LoginAck     , MT_Blank        , On_LoginAck     , "player login ack from client"},
	MT_LogoutReq     :&NetMsgCb{MT_LogoutReq    , MT_LogoutAck    , On_LogoutReq    , ""},
	MT_LogoutAck     :&NetMsgCb{MT_LogoutAck    , MT_Blank        , On_LogoutAck    , ""},
	MT_LoginRoomReq  :&NetMsgCb{MT_LoginRoomReq , MT_LoginRoomAck , On_LoginRoomReq , ""},
	MT_LoginRoomAck  :&NetMsgCb{MT_LoginRoomAck , MT_Blank        , On_LoginRoomAck , ""},
	MT_ScreenChangeNotify :&NetMsgCb{MT_ScreenChangeNotify, MT_ScreenChangeNotify, On_ScreenChangeNotify, ""},
	MT_LeaveRoomReq  :&NetMsgCb{MT_LeaveRoomReq , MT_LeaveRoomAck , On_LeaveRoomReq , ""},
	MT_LeaveRoomAck  :&NetMsgCb{MT_LeaveRoomAck , MT_Blank        , On_LeaveRoomAck , ""},
}
