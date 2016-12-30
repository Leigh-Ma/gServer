package play

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"library/logger"
	. "types"
)

type Player struct {
	UserId   IdString
	Name     string
	UUID     string
	HeroName string
	Skin     string
	Skills   []string
	room     *Room
}

func (p *Player) FillActiveDetail(sa *ActDetail) {
	sa.Name = p.Name
	sa.SubType = p.HeroName
	sa.Skin = p.Skin

	//sa.Hp = p.Hp
	//sa.FullHp = p.FullHp
}

func (p *Player) UpSert(session *mgo.Session) bool {
	c := session.Clone().DB(dbName).C(MONGO_COLLECTION_PLAYERS)
	if _, err := c.Upsert(bson.M{"UserId": string(p.UserId)}, p); err != nil {
		logger.Error("Player UpSertDb error: %s", err.Error())
		return false
	}
	return true
}

func (p *Player) Destroy(session *mgo.Session) bool {
	c := session.Clone().DB(dbName).C(MONGO_COLLECTION_PLAYERS)
	if err := c.Remove(bson.M{"UserId": string(p.UserId)}); err != nil {
		logger.Error("Player UpSertDb error: %s", err.Error())
		return false
	}
	return true
}
