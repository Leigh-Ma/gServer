package play

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"library/logger"
	"sync"
	. "types"
)

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

func (pm *PlayerManager) GetPlayer(userId IdString) (*Player, bool) {
	pm.Lock()
	p, ok := pm.Players[userId]
	pm.Unlock()
	return p, ok
}

func (pm *PlayerManager) AddPlayer(p *Player) {
	pm.Lock()
	pm.Players[p.UserId] = p
	pm.Unlock()
}

func (pm *PlayerManager) DelPlayer(userId IdString) {
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

func (pm *PlayerManager) LoadOneFromDb(session *mgo.Session, userId IdString) (*Player, bool) {
	if session == nil {
		return nil, false
	}

	var player *Player
	c := session.Clone().DB(MONGO_DB_NAME).C(MONGO_COLLECTION_PLAYERS)
	err := c.Find(bson.M{"userid": userId}).One(&player)
	if err != nil {
		logger.Info("Load Player %s From DataBase error: %s", userId, err.Error())
		return nil, false
	}

	if player == nil {
		return nil, false
	}

	pm.AddPlayer(player)

	return player, true
}
