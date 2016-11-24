package protohandle

import (
	pb "github.com/golang/protobuf/proto"
	dm "library/core/datamsg"
	"library/logger"
	"play"
	. "types"
)

func PlayHandler(payload *NetMsg) bool {
	play.PlayFrame.SetFrameTime()
	logger.Info("play frame time set to %d", play.PlayFrame.FrameTime())
	handler, ok := play.NetMsgTypeHandler[payload.OpCode]
	if !ok {
		logger.Info("recv invalid payload type %d", payload.OpCode)
		return false
	}

	logger.Info("recv payload type %d", payload.OpCode)

	ack := handler.Handler(payload.UserId.ToIdString(), payload.OpCode, payload.Content)
	if ack != nil {
		pbAck, ok := ack.(pb.Message)
		if !ok {
			logger.Info("invalid ack payload for req %d", payload.OpCode)
			return false
		}
		payload.OpCode = handler.RetCode
		payload.SetPayLoad(pbAck)
		logger.Info("ack payload type %d", payload.OpCode)
	}

	return true
}

func (t *protoHandle) DataHandler(msg *dm.DataMsg) bool {
	payload, ok := msg.Payload.(*NetMsg)
	if !ok {
		logger.Info("recv invalid message type %d", msg.MsgType)
		return false
	}

	if ok = PlayHandler(payload); ok {
		msg.Payload, _ = payload.BinaryProtoNoId()
		msg.Receiver = msg.Sender
		msg.Sender = ServiceName
	}

	if msg.Receiver != dm.NoReceiver {
		t.Output.WritePipeNoBlock(msg)
	}
	return true
}
