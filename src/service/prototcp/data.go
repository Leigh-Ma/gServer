package prototcp

import (
	dm "library/core/datamsg"
	"library/logger"
	"net"
)

func (t *protoTCPServer) DataHandler(msg *dm.DataMsg) bool {
	meta, ok := msg.Meta(t.Name)
	if !ok {
		logger.Error("%s:wrong meta in datamsg:%+v", t.Name, msg)
		return false
	}
	connection := meta.(*net.TCPConn)
	//todo: need to verify if the data payload is []byte
	content, ok := msg.Payload.([]byte)
	if !ok {
		logger.Warn("payload is not in form []byte, err:%s", t.Name)
	}
	count, err := connection.Write(content)
	if err != nil {
		logger.Warn("%s:conn write err:%s", t.Name, err.Error())
		connection.Close()
		return false
	} else {
		logger.Info("%s:sent to network:%d byte", t.Name, count)
	}
	return true
}
