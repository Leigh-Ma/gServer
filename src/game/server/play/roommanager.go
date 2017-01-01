package play

import (
	"sync"
	"types"
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

func (rm *RoomManager) CreateRoom(id types.IdString, name string) *Room {
	room := NewRoom(id, name)
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
	for _, r := range rm.Rooms {
		if r.BcgMemberNum() < maxRoomMemberNum {
			room = r
			break
		}
	}
	return
}
