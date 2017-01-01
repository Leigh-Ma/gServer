package gatewayinner

import (
	"game/gateway/gm"
	"library/logger"
	. "types"
)

//client req message, forward to server with userId
func GatewayOnClientMessage(clientMeta *gm.ConnMeta, msg *NetMsg) bool {
	logger.Info("gateway: MSG <%16s> from player <%s> received", msg.TypeString(), clientMeta.ID)

	serverMeta := clientMeta.ForwardMeta

	code := msg.Code()
	switch code {
	case MT_LoginReq:
		ok := gm.Clients.Login(clientMeta, msg.Content)
		if !ok {
			logger.Error("gateway: player %s, login server error", clientMeta.ID)
			return false
		}
		user, _ := gm.UserM.GetByUserId(clientMeta.ID)

		req := &LoginReq{UserId: string(user.UserId), Uuid: user.Uuid, ServerId: string(user.ServerId)}
		msg.SetPayLoad(MT_LoginReq, req, NetMsgIdFlagClient)

		serverMeta = clientMeta.ForwardMeta
	case MT_LogoutReq:
		serverMeta = clientMeta.ForwardMeta
		gm.Clients.Logout(clientMeta)
	case MT_SearchRoomReq:
		roomId, ok := gm.Clients.SearchRoomServer(clientMeta, msg.Content)
		user, _ := gm.UserM.GetByUserId(clientMeta.ID)
		if ok {
			req := &SearchRoomReq{Tag: string(roomId), ServerId: string(user.ServerId)}
			msg.SetPayLoad(MT_LoginReq, req, NetMsgIdFlagClient)
		}
		serverMeta = clientMeta.ForwardMeta
	}

	if serverMeta == nil {
		logger.Error("gateway: player %s, server not distributed", clientMeta.ID)
		return false
	}

	msg.ObjectID = clientMeta.ObjID

	binary, err := msg.BinaryProto()
	if err != nil {
		logger.Error("gateway: player %s, payload marshal error: %s", clientMeta.ID, err.Error())
		return false
	}

	return serverMeta.CsToServer(msg.TypeString(), binary)
}
