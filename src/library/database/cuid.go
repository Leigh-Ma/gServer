package database

import (
	"fmt"
	"gopkg.in/mgo.v2"
	dm "library/core/datamsg"
	"library/logger"
	"library/structenh"
	"service"
)

const (
	dbOperateUpdate = int8(1)
	dbOperateDelete = int8(2)
)

var operateTypeName = map[int8]string{
	dbOperateUpdate: "dbOperateUpdate",
	dbOperateDelete: "dbOperateDelete",
}

type IDbOperate interface {
	UpSert(session *mgo.Session) bool
	Destroy(session *mgo.Session) bool
}

type dbOperate struct {
	object  IDbOperate
	operate int8
}

func (op *dbOperate) CRUD(session *mgo.Session) bool {
	switch op.operate {
	case dbOperateUpdate:
		return op.object.UpSert(session)
	case dbOperateDelete:
		return op.object.Destroy(session)
	default:
		logger.Error("Database operation failed: invalid operate type %d", op.operate)
		return false
	}
}

func (op *dbOperate) Inspect() string {
	opName, ok := operateTypeName[op.operate]
	if !ok {
		opName = "dbOperateInvalid"
	}

	return fmt.Sprintf("%s: %s", opName, structenh.Stringify(op.object))
}

func DbUpdate(dbObj IDbOperate) {
	service.ServicePool.SendData(dm.NewDataMsg("", "mongo", 0, &dbOperate{object: dbObj, operate: dbOperateUpdate}))
}

func DbDestroy(dbObj IDbOperate) {
	service.ServicePool.SendData(dm.NewDataMsg("", "mongo", 0, &dbOperate{object: dbObj, operate: dbOperateDelete}))
}
