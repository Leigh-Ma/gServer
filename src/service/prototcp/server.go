package prototcp

import (
	"net"
)

import (
	dm "library/core/datamsg"
	"service"
)

const ServiceName = "prototcpserver"

type protoTCPServer struct {
	*service.Service
	output *dm.DataMsgPipe

	listener *net.TCPListener
	ip       string
	port     string
}

func NewProtoTCPServer(name, ip, port string) *protoTCPServer {
	t := &protoTCPServer{}
	t.Service = service.NewService(ServiceName)
	t.output = nil
	t.Name = name
	t.ip = ip
	t.port = port
	t.Instance = t
	return t
}
