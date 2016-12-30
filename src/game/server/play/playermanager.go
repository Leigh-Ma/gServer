package play

import (
	"sync"
	. "types"
)

var OnlineM = NewPlayersManager()

type PlayerManager struct {
	Players map[IdString]*Player
	sync.Mutex
}

func NewPlayersManager() *PlayerManager {
	return &PlayerManager{
		Players: make(map[IdString]*Player, 0),
	}
}

func (pm *PlayerManager) GetOnlinePlayer(userId IdString) (*Player, bool) {
	pm.Lock()
	p, ok := pm.Players[userId]
	pm.Unlock()
	return p, ok
}

func (pm *PlayerManager) AddOnlinePlayer(p *Player) {
	pm.Lock()
	pm.Players[p.UserId] = p
	pm.Unlock()
}

func (pm *PlayerManager) DelOnlinePlayer(userId IdString) {
	pm.Lock()
	delete(pm.Players, userId)
	pm.Unlock()
}