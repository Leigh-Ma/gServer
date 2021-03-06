package gatewayoutter

import (
	"game/com"
	"game/gateway/gm"
	pb "github.com/golang/protobuf/proto"
	"library/logger"
	"library/structenh"
	"netmsghandle/gs"
	. "types"
)

func forwardServerMessageToClient(serverMeta *gm.ConnMeta, msg *NetMsg) bool {
	go func() {
		opName := msg.TypeString()

		clientMeta, ok := gm.Clients.GetMeta(msg.ToIdString())
		if !ok || clientMeta.Conn == nil {
			logger.Error("cs: MSG <%16s>: player %s connection not on", opName, msg.ToIdString())
			return
		}

		binary, err := msg.BinaryProtoToClient()
		if err != nil {
			logger.Error("cs: MSG <%16s>: payload marshal error %s", opName, err.Error())
			return
		}

		clientMeta.CsToClient(opName, binary)
		return
	}()
	return true
}

//gateway receives message from server (protocol between server & gateway)
func handleServer2GatewayMessage(serverMeta *gm.ConnMeta, msg *NetMsg) bool {
	opName := msg.TypeString()
	if msg.Code() == MT_ServerLoginReq {
		gm.Servers.Login(serverMeta, msg.Content)
	}

	if msg.ObjectID != serverMeta.ID.ToObjectID() {
		logger.Error("gs: MSG <%16s>: meta server_id %s != %s in message head", opName, serverMeta.ID, msg.ToIdString())
		return false
	}

	handler, ok := gs.NetMsgTypeHandler[msg.Code()]
	if !ok {
		logger.Error("gs: MSG <%16s>: server %s, handler not exist", opName, msg.ToIdString())
		return false
	}

	ack := handler.Handler(msg.ToIdString(), msg.Code(), msg.Content)
	if ack == nil {
		logger.Report("gs: MSG <%16s>: ACK <%16s> nil, ", opName, handler.RetCode.TypeString())
		return true
	}

	pbAck, ok := ack.(pb.Message)
	if !ok {
		logger.Error("gs: MSG <%16s>: ACK <%16s> payload invalid", opName, handler.RetCode.TypeString())
		return false
	}

	msg.SetPayLoad(handler.RetCode, pbAck, NetMsgIdFlagServer)

	bin, err := msg.BinaryProto()
	if err != nil {
		logger.Error("gs: MSG <%16s>: payload marshal error %s", msg.TypeString(), err.Error())
		return false
	}

	logger.Payload("rx-x: %4d: payload %s", handler.RetCode, structenh.StringifyStruct(ack))
	return serverMeta.GsToServer(opName, msg.TypeString(), bin)
}

func distributeBroadCastMessageToClients(serverMeta *gm.ConnMeta, msg *NetMsg) bool {
	go func() {
		opName := msg.TypeString()
		grp, ok := com.FindBrdCastGroup(serverMeta.ID, msg.ObjectID.ToIdString())
		if !ok {
			logger.Error("brdcast: MSG: <%16s>, group %s not found", opName, msg.ToIdString())
			return
		}

		binary, err := msg.BinaryProtoToClient()
		if err != nil {
			logger.Error("brdcast: MSG: <%16s>, payload marshal error %s", opName, err.Error())
			return
		}

		for playerId := range grp.Members {
			clientMeta, ok := gm.Clients.GetMeta(playerId)
			if !ok || clientMeta.Conn == nil {
				logger.Error("brdcast: MSG <%16s>, player %s meta not valid", opName, playerId)
				continue
			}
			clientMeta.BroadCastSendClient(opName, binary)
		}
	}()

	return true
}

//server ack/heartbeat message, forward to client without userId
func HandleMessageFromServer(serverMeta *gm.ConnMeta, msg *NetMsg) bool {
	//server message for
	switch {
	case msg.HasFlag(NetMsgIdFlagClient):
		return forwardServerMessageToClient(serverMeta, msg)
	case msg.HasFlag(NetMsgIdFlagServer):
		return handleServer2GatewayMessage(serverMeta, msg)
	case msg.HasFlag(NetMsgIdFlagBroadCast):
		return distributeBroadCastMessageToClients(serverMeta, msg)
	}

	logger.Info("gateway: server message with invalid flag, code %d",
		msg.TypeString())

	return false
}
