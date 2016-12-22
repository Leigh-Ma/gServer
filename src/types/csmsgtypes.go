package types

//Auto generated, do not modify unless you know clearly what you are doing.

const (
    MT_LoginReq          =  MsgType(  1)
    MT_LoginAck          =  MsgType(  2)
    MT_LogoutReq         =  MsgType(  3)
    MT_LogoutAck         =  MsgType(  4)
    MT_LoginRoomReq      =  MsgType(  8)
    MT_LoginRoomAck      =  MsgType(  9)
    MT_ScreenChangeNotify =  MsgType( 11)
    MT_LeaveRoomReq      =  MsgType( 12)
    MT_LeaveRoomAck      =  MsgType( 13)
)

var csmsgtypesNames = map[MsgType]string {
    MT_Blank            :  "MT_Blank",
    MT_LoginReq         :  "MT_LoginReq",
    MT_LoginAck         :  "MT_LoginAck",
    MT_LogoutReq        :  "MT_LogoutReq",
    MT_LogoutAck        :  "MT_LogoutAck",
    MT_LoginRoomReq     :  "MT_LoginRoomReq",
    MT_LoginRoomAck     :  "MT_LoginRoomAck",
    MT_ScreenChangeNotify:  "MT_ScreenChangeNotify",
    MT_LeaveRoomReq     :  "MT_LeaveRoomReq",
    MT_LeaveRoomAck     :  "MT_LeaveRoomAck",
}
