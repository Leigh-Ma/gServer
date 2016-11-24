package prototcp

import (
	"io"
	cm "library/core/controlmsg"
	dm "library/core/datamsg"
	"library/logger"
	"net"
	"service/protohandle"
	"time"
	. "types"
)

func (t *protoTCPServer) Start(bus *dm.DataMsgPipe) bool {
	t.output = bus
	tcpAddr, err := net.ResolveTCPAddr("tcp", t.ip+":"+t.port)
	if err != nil {
		logger.Error("%s:net.ResolveTCPAddr error,%s", t.Name, err.Error())
		return false
	}

	t.listener, err = net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		logger.Error("%s:net.ListenTCP error,%s", t.Name, err.Error())
		return false
	}

	logger.Info("%s:listening port:%s", t.Name, t.port)
	go t.serve()
	return true
}

func (t *protoTCPServer) Pause() bool {
	return true
}

func (t *protoTCPServer) Resume() bool {
	return true
}

func (t *protoTCPServer) Exit() bool {
	return true
}

func (t *protoTCPServer) ControlHandler(msg *cm.ControlMsg) (int, int) {
	return cm.NextActionContinue, cm.ProcessStatOK
}

func (t *protoTCPServer) serve() {

	for {
		connect, err := t.listener.AcceptTCP()
		if err != nil {
			logger.Error("%s:listener.AcceptTCP error,%s", t.Name, err.Error())
			time.Sleep(time.Second * 2)
			connect.Close()
			continue
		}
		go t.readConn(connect)
	}
}

func (t *protoTCPServer) readConn(connection *net.TCPConn) {

	for {
		head := make([]byte, NetMsgHeadSize)
		n, err := io.ReadAtLeast(connection, head, NetMsgHeadSize)
		if err != nil {
			logger.Warn("%s:read byte:%d,error:%s", t.Name, n, err.Error())
			if err == io.EOF {
				return
			}
			connection.Close()
			return
		}
		msg, err := NewNetMsgFromHead(head)
		if err != nil {
			logger.Warn("%s:decode msg head error:%s", t.Name, n, err.Error())
			if err == io.EOF {
				return
			}
			connection.Close()
			return
		}
		_, err = io.ReadAtLeast(connection, msg.Content, int(msg.Size))
		if err != nil {
			logger.Warn("%s:read byte:%d,error:%s", t.Name, n, err.Error())
			if err == io.EOF {
				return
			}
			connection.Close()
			return
		}

		d := dm.NewDataMsg(ServiceName, protohandle.ServiceName, MsgTypeProtoNet, msg)
		d.SetMeta(t.Name, connection)
		t.output.WritePipeNoBlock(d)
	}
}
