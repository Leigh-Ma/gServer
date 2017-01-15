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
	MT_SetPlayerNameReq :&NetMsgCb{MT_SetPlayerNameReq, MT_SetPlayerNameAck, On_SetPlayerNameReq, ""},
	MT_SetPlayerNameAck :&NetMsgCb{MT_SetPlayerNameAck, MT_Blank        , On_SetPlayerNameAck, ""},
	MT_ChoseSideReq  :&NetMsgCb{MT_ChoseSideReq , MT_ChoseSideAck , On_ChoseSideReq , ""},
	MT_ChoseSideAck  :&NetMsgCb{MT_ChoseSideAck , MT_Blank        , On_ChoseSideAck , ""},
	MT_StartFightReq :&NetMsgCb{MT_StartFightReq, MT_StartFightAck, On_StartFightReq, ""},
	MT_StartFightAck :&NetMsgCb{MT_StartFightAck, MT_Blank        , On_StartFightAck, ""},
	MT_LeaveRoomReq  :&NetMsgCb{MT_LeaveRoomReq , MT_LeaveRoomAck , On_LeaveRoomReq , ""},
	MT_LeaveRoomAck  :&NetMsgCb{MT_LeaveRoomAck , MT_Blank        , On_LeaveRoomAck , ""},
	MT_MoveActionReq :&NetMsgCb{MT_MoveActionReq, MT_MoveActionAck, On_MoveActionReq, ""},
	MT_MoveActionAck :&NetMsgCb{MT_MoveActionAck, MT_Blank        , On_MoveActionAck, ""},
	MT_ChoseHeroReq  :&NetMsgCb{MT_ChoseHeroReq , MT_ChoseHeroAck , On_ChoseHeroReq , ""},
	MT_ChoseHeroAck  :&NetMsgCb{MT_ChoseHeroAck , MT_Blank        , On_ChoseHeroAck , ""},
	MT_SearchRoomReq :&NetMsgCb{MT_SearchRoomReq, MT_SearchRoomAck, On_SearchRoomReq, ""},
	MT_SearchRoomAck :&NetMsgCb{MT_SearchRoomAck, MT_Blank        , On_SearchRoomAck, ""},
	MT_ScreenChangeNotify :&NetMsgCb{MT_ScreenChangeNotify, MT_ScreenChangeNotify, On_ScreenChangeNotify, ""},
	MT_StartFightNotify :&NetMsgCb{MT_StartFightNotify, MT_StartFightNotify, On_StartFightNotify, ""},
	MT_ChoseSideNotify :&NetMsgCb{MT_ChoseSideNotify, MT_ChoseSideNotify, On_ChoseSideNotify, ""},
}
