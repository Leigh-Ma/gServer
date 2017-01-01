package gm

import (
	"game/com"
	"library/frame"
	"library/logger"
	. "types"
)

var GatewayFrame = frame.NewFrame()

func serverCommonAck(code int32) *ServerCommonAck {
	return &ServerCommonAck{ErrCode: code}
}

func choseServerToLogin(client *ConnMeta, user *User) (IdString, bool) {
	roomId, serverId := InvalidIdString, InvalidIdString
	serverMeta := client.ForwardMeta
	if serverMeta != nil && user.ServerId == serverMeta.ID && user.ServerId.IsValid() {
		return roomId, true
	}

	serverMeta, ok := Servers.GetMeta(user.ServerId)
	if !ok || serverMeta.Conn == nil {
		serverId, roomId = com.ChoseARoom()
		serverMeta, _ = Servers.GetMeta(serverId)
	}

	if serverMeta == nil {
		logger.Info("no suitable server for client %s", client.ID)
		return roomId, false
	}

	logger.Info("client %s distribute to server %s", client.ID, serverMeta.ID)
	client.ForwardMeta = serverMeta
	user.ServerId = serverMeta.ID

	return roomId, true
}
