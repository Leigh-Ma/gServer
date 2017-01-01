package gm

import (
	"library/frame"
	"library/logger"
	. "types"
)

var GatewayFrame = frame.NewFrame()

func serverCommonAck(code int32) *ServerCommonAck {
	return &ServerCommonAck{ErrCode: code}
}

//todo chose according to room information
func choseServerToLogin(client *ConnMeta, user *User) bool {
	serverMeta, ok := Servers.GetMeta(user.ServerId)
	if !ok || serverMeta.Conn == nil {
		for _, meta := range Servers.connections {
			serverMeta = meta
			break
		}
	}

	if serverMeta == nil {
		logger.Info("no suitable server for client %s", client.ID)
		return false
	}

	logger.Info("client %s distribute to server %s", client.ID, serverMeta.ID)
	client.ForwardMeta = serverMeta
	user.ServerId = serverMeta.ID

	return true
}
