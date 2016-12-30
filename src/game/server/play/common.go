package play

import (
	"game/server/support"
	"library/config"
	"library/frame"
	. "types"
)

var dbName string

func init() {
	dbName, _ = config.Cfg.GetString("server_name")
}

var AsyncSender = new(support.AsyncSender)

var PlayFrame = frame.NewFrame()

func getCommonAck(code int32) *CommonAck {
	return &CommonAck{
		Status:    code,
		Message:   ErrDesc[code],
		TimeStamp: uint32(PlayFrame.FrameTime()),
	}
}

func getMyServerId() IdString {
	return support.MyServerId
}
