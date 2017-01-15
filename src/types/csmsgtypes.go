package types

//Auto generated, do not modify unless you know clearly what you are doing.

const (
    MT_LoginReq          =  MsgType(  1)
    MT_LoginAck          =  MsgType(  2)
    MT_LogoutReq         =  MsgType(  3)
    MT_LogoutAck         =  MsgType(  4)
    MT_SetPlayerNameReq  =  MsgType(  5)
    MT_SetPlayerNameAck  =  MsgType(  6)
    MT_ChoseSideReq      =  MsgType(  7)
    MT_ChoseSideAck      =  MsgType(  8)
    MT_StartFightReq     =  MsgType(  9)
    MT_StartFightAck     =  MsgType( 10)
    MT_LeaveRoomReq      =  MsgType( 11)
    MT_LeaveRoomAck      =  MsgType( 12)
    MT_MoveActionReq     =  MsgType( 13)
    MT_MoveActionAck     =  MsgType( 14)
    MT_ChoseHeroReq      =  MsgType( 15)
    MT_ChoseHeroAck      =  MsgType( 16)
    MT_SearchRoomReq     =  MsgType( 17)
    MT_SearchRoomAck     =  MsgType( 18)
    MT_ScreenChangeNotify =  MsgType(101)
    MT_StartFightNotify  =  MsgType(102)
)

var csmsgtypesNames = map[MsgType]string {
    MT_Blank            :  "MT_Blank",
    MT_LoginReq         :  "MT_LoginReq",
    MT_LoginAck         :  "MT_LoginAck",
    MT_LogoutReq        :  "MT_LogoutReq",
    MT_LogoutAck        :  "MT_LogoutAck",
    MT_SetPlayerNameReq :  "MT_SetPlayerNameReq",
    MT_SetPlayerNameAck :  "MT_SetPlayerNameAck",
    MT_ChoseSideReq     :  "MT_ChoseSideReq",
    MT_ChoseSideAck     :  "MT_ChoseSideAck",
    MT_StartFightReq    :  "MT_StartFightReq",
    MT_StartFightAck    :  "MT_StartFightAck",
    MT_LeaveRoomReq     :  "MT_LeaveRoomReq",
    MT_LeaveRoomAck     :  "MT_LeaveRoomAck",
    MT_MoveActionReq    :  "MT_MoveActionReq",
    MT_MoveActionAck    :  "MT_MoveActionAck",
    MT_ChoseHeroReq     :  "MT_ChoseHeroReq",
    MT_ChoseHeroAck     :  "MT_ChoseHeroAck",
    MT_SearchRoomReq    :  "MT_SearchRoomReq",
    MT_SearchRoomAck    :  "MT_SearchRoomAck",
    MT_ScreenChangeNotify:  "MT_ScreenChangeNotify",
    MT_StartFightNotify :  "MT_StartFightNotify",
}
