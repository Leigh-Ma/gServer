package types

import "fmt"

//unix timestamp
type UnixTS int64

type IdString string

type PlayerId uint64

func (p PlayerId) ToIdString() IdString {
	return IdString(fmt.Sprintf("%x", p))
}

const (
	MsgTypeUnknown  = 0
	MsgTypeTelnet   = 1
	MsgTypeProtoNet = 2
	MsgTypeJsonSim  = 3
)