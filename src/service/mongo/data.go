package mongo

import (
//"time"
)

import (
	dm "library/core/datamsg"
	"library/logger"
)

func (t *mongoType) DataHandler(msg *dm.DataMsg) bool {
	//logger.Info("this is handle:%+v", msg)
	d, ok := msg.Payload.(Dirty)
	if !ok {
		logger.Error("mongo: get Dirty interface error")
		return false
	}

	for ok = d.CRUD(t.session); !ok; {
		logger.Info("CRUD failed:%s", d.Inspect())
		ok = d.CRUD(t.session)
	}

	if msg.Sender != "" {
		msg.Receiver = msg.Sender
		msg.Sender = t.Name
		t.output.WritePipeNoBlock(msg)
	}
	return true
}
