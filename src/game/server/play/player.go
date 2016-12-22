package play

import "types"

type Player struct {
	UserId    types.IdString
	SubType   int8
	Skin      int8
	Direction int8

	Name   string
	Hp     int32
	FullHp int32

	room *Room
}
