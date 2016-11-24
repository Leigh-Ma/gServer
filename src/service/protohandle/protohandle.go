package protohandle

import (
	dm "library/core/datamsg"
	"service"
)

const ServiceName = "protohandle"

type protoHandle struct {
	*service.Service
	Output *dm.DataMsgPipe
}

func NewProtoHandle(name string) *protoHandle {
	t := &protoHandle{}
	t.Service = service.NewService(ServiceName)
	t.Output = nil
	t.Name = name
	t.Instance = t
	return t
}
