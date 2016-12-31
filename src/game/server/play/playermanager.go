package play

import (
	"gopkg.in/mgo.v2"
	"library/logger"
	"sync"
	. "types"
)

var OnlineM = NewPlayersManager()
var AllPlayerM = NewPlayersManager()

type PlayerManager struct {
	Players map[IdString]*Player
	sync.Mutex
}

func NewPlayersManager() *PlayerManager {
	return &PlayerManager{
		Players: make(map[IdString]*Player, 0),
	}
}

func (pm *PlayerManager) GetOnePlayer(userId IdString) (*Player, bool) {
	pm.Lock()
	p, ok := pm.Players[userId]
	pm.Unlock()
	return p, ok
}

func (pm *PlayerManager) AddOnePlayer(p *Player) {
	pm.Lock()
	pm.Players[p.UserId] = p
	pm.Unlock()
}

func (pm *PlayerManager) DelOnePlayer(userId IdString) {
	pm.Lock()
	delete(pm.Players, userId)
	pm.Unlock()
}

func (pm *PlayerManager) LoadAllFrommDb(session *mgo.Session) bool {
	c := session.Clone().DB(MONGO_DB_NAME).C(MONGO_COLLECTION_PLAYERS)
	remain, err := c.Count()
	total := remain

	if err != nil {
		logger.Error("Load Players From DataBase read count error: %s", err.Error())
		return false
	}

	cursor, step := 0, 1000
	for {
		if remain <= 0 {
			break
		}

		if remain < step {
			step = remain
		}

		players := make([]*Player, 0, step)
		err = c.Find(nil).Skip(cursor).Limit(step).All(&players)
		if err != nil {
			logger.Error("Load User From DataBase get data error: %s", err.Error())
			return false
		}

		for _, player := range players {
			pm.Players[player.UserId] = player
		}

		cursor += step
		remain -= step
	}

	logger.Info("Load %d Players From DataBase", total)
	return true
}
