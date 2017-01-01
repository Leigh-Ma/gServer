package database

import "gopkg.in/mgo.v2"

var MongoProxy *mgo.Session

func SetMongoProxy(session *mgo.Session) {
	MongoProxy = session
}
