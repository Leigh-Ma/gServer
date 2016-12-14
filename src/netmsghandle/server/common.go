package server

import (
	"library/frame"
	. "types"
)

var ServerFrame = frame.NewFrame()

func getCommonAck(code int32) *CommonAck {
	return &CommonAck{
		Status:    code,
		Message:   ErrDesc[code],
		TimeStamp: uint32(ServerFrame.FrameTime()),
	}
}
