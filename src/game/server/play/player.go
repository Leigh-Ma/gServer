package play

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"library/logger"
	. "types"
)

type Player struct {
	UserId IdString
	UUID   string
	Hero
	room *Room
}

func (p *Player) RoleInfo() *RoleInfo {
	hi := &RoleInfo{
		UserId:   string(p.UserId),
		UserName: p.UserName,
		HeroType: p.HeroType,
		FullHp:   p.FullHp,
		Hp:       p.Hp,
		Skin:     p.Skin,
	}
	hi.Skills = append(hi.Skills, p.Skills...)

	return hi
}

func (p *Player) FillActiveDetail(sa *ActDetail) {
	sa.Name = p.UserName
	sa.SubType = p.HeroType
	sa.Skin = p.Skin

	sa.Hp = p.Hp
	sa.FullHp = p.FullHp
	sa.Skills = make([]string, len(p.Skills))
	copy(sa.Skills, p.Skills)
}

func (p *Player) Upsert(session *mgo.Session) bool {
	c := session.Clone().DB(MONGO_DB_NAME).C(MONGO_COLLECTION_PLAYERS)
	if _, err := c.Upsert(bson.M{"userid": string(p.UserId)}, p); err != nil {
		logger.Error("Player Upsert error: %s", err.Error())
		return false
	}
	return true
}

func (p *Player) Destroy(session *mgo.Session) bool {
	c := session.Clone().DB(MONGO_DB_NAME).C(MONGO_COLLECTION_PLAYERS)
	if err := c.Remove(bson.M{"userid": string(p.UserId)}); err != nil {
		logger.Error("Player Destroy error: %s", err.Error())
		return false
	}
	return true
}
