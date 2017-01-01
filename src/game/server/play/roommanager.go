package play

import (
	"library/idgen"
	"sync"
	"types"
	"game/com"
)

var RoomM = NewRoomManager()

type RoomManager struct {
	Rooms map[types.IdString]*Room
	sync.Mutex
}

func NewRoomManager() *RoomManager {
	return &RoomManager{
		Rooms: make(map[types.IdString]*Room, 0),
	}
}

func (rm *RoomManager) FindRoom(roomId types.IdString) (*Room, bool) {
	rm.Lock()
	room, ok := rm.Rooms[roomId]
	rm.Unlock()
	return room, ok
}

func (rm *RoomManager) CreateRoom(name string) *Room {
	room := NewRoom(idgen.NewObjectID().ToIdString(), name)
	rm.Lock()
	rm.Rooms[room.Id] = room
	rm.Unlock()
	return room
}

func (rm *RoomManager) DestroyRoom(room *Room) {
	rm.Lock()
	delete(rm.Rooms, room.Id)
	rm.Unlock()
	room.Destroy()
}

func (rm *RoomManager) ChoseByTag(tag string) (room *Room) {
	room, ok := rm.FindRoom(types.IdString(tag))
	if ok && room.BcgMemberNum() < maxRoomMemberNum {
		return room
	}

	_, roomId := com.ChoseARoom()
	room, _ = rm.FindRoom(roomId)
	return room
}
