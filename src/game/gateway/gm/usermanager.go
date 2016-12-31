package gm

import (
	"gopkg.in/mgo.v2"
	"library/idgen"
	"library/logger"
	. "types"
)

var UserM = NewUserManager()

type User struct {
	Uuid     string
	UserId   IdString
	ServerId IdString
}

func NewUser(uuid string) *User {
	return &User{
		Uuid:   uuid,
		UserId: idgen.NewObjectID().ToIdString(),
	}
}

type UserManager struct {
	UsersByUuid   map[string]*User
	UsersByUserId map[IdString]*User
}

func NewUserManager() *UserManager {
	return &UserManager{
		UsersByUuid:   make(map[string]*User, 0),
		UsersByUserId: make(map[IdString]*User, 0),
	}
}

func (um *UserManager) LoadAllFrommDb(session *mgo.Session) {
	c := session.Clone().DB(MONGO_DB_NAME).C(MONGO_COLLECTION_PLAYERS)
	remain, err := c.Count()

	if err != nil {
		logger.Error("Load User From DataBase read count error: %s", err.Error())
		return
	}

	cursor, step := 0, 1000
	for {
		if remain <= 0 {
			return
		}

		if remain < step {
			step = remain
		}

		users := make([]*User, 0, step)
		err = c.Find(nil).Skip(cursor).Limit(step).All(&users)
		if err != nil {
			logger.Error("Load User From DataBase get data error: %s", err.Error())
			return
		}

		for _, user := range users {
			um.UsersByUuid[user.Uuid] = user
			um.UsersByUserId[user.UserId] = user
		}

		cursor += step
		remain -= step
	}

}

func (um *UserManager) GetByUserId(userId IdString) (*User, bool) {
	user, ok := um.UsersByUserId[userId]
	return user, ok
}

func (um *UserManager) GetByUuid(uuid string) (*User, bool) {
	user, ok := um.UsersByUuid[uuid]
	return user, ok
}

func (um *UserManager) CreateNewUser(uuid string) *User {
	user := NewUser(uuid)
	um.UsersByUserId[user.UserId] = user
	um.UsersByUuid[user.Uuid] = user

	return user
}
