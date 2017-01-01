package gm

import (
	"errors"
	pb "github.com/golang/protobuf/proto"
	"library/logger"
	"net"
	"sync"
	. "types"
)

const (
	connMetaTypeServer = 1
	connMetaTypeClient = 2
)

var Servers = &metaManage{
	connections: make(map[IdString]*ConnMeta, 0),
	metaType:    connMetaTypeServer,
}

var Clients = &metaManage{
	connections: make(map[IdString]*ConnMeta, 0),
	metaType:    connMetaTypeClient,
}

func NewServerMeta(conn *net.TCPConn) *ConnMeta {
	return &ConnMeta{metaType: connMetaTypeServer, Conn: conn}
}

func NewClientMeta(conn *net.TCPConn) *ConnMeta {
	return &ConnMeta{metaType: connMetaTypeClient, Conn: conn}
}

type ConnMeta struct {
	ID          IdString
	Conn        *net.TCPConn
	LoginAt     UnixTS
	LastSend    UnixTS
	LastRecv    UnixTS
	SentByte    int
	RecvByte    int
	SentNum     int
	RecvNum     int
	ForwardMeta *ConnMeta

	ObjID ObjectID

	metaType int
}

type metaManage struct {
	lock        sync.Mutex
	connections map[IdString]*ConnMeta
	metaType    int
}

func (mm *metaManage) mineMeta(meta *ConnMeta) bool {
	if meta == nil {
		return false
	}
	if meta.metaType != mm.metaType {
		logger.Error("Connection meta useage error: mm %d accept %d", mm.metaType, meta.metaType)
		return false
	}
	return true
}

func (mm *metaManage) GetMeta(id IdString) (*ConnMeta, bool) {
	mm.lock.Lock()
	meta, ok := mm.connections[id]
	mm.lock.Unlock()
	return meta, ok
}

func (mm *metaManage) Logout(meta *ConnMeta) {
	if ok := mm.mineMeta(meta); !ok {
		return
	}

	mm.lock.Lock()
	delete(mm.connections, meta.ID)
	mm.lock.Unlock()

	if meta.Conn != nil {
		meta.Conn.Close()
		meta.Conn = nil
	}
	meta.ForwardMeta = nil
}

func (mm *metaManage) userLogin(meta *ConnMeta, content []byte) bool {
	var (
		user *User
		ok   bool
	)

	req := &LoginReq{}
	if err := pb.Unmarshal(content, req); err != nil {
		return false
	}

	if req.Uuid == "" {
		logger.Warn("gateway: login request uuid is blank!!")
		return false
	}

	if user, ok = UserM.GetByUserId(IdString(req.UserId)); user == nil {
		user, ok = UserM.GetByUuid(req.Uuid)
	}

	if !ok {
		user = UserM.CreateNewUser(req.Uuid)
		logger.Info("gateway: client uuid %s not exist, create new: %v", req.Uuid, *user)
	}

	meta.ID = user.UserId
	user.ServerId = IdString(req.ServerId)

	return choseServerToLogin(meta, user)
}

func (mm *metaManage) serverLogin(meta *ConnMeta, content []byte) bool {
	req := &ServerLoginReq{}
	if err := pb.Unmarshal(content, req); err != nil {
		return false
	}
	meta.ID = IdString(req.ServerId)
	return true
}

func (mm *metaManage) Login(meta *ConnMeta, content []byte) bool {
	ok := mm.mineMeta(meta)
	if !ok {
		return false
	}

	switch meta.metaType {
	case connMetaTypeClient:
		ok = mm.userLogin(meta, content)

	case connMetaTypeServer:
		ok = mm.serverLogin(meta, content)
	}
	if !ok {
		return false
	}

	meta.ObjID = meta.ID.ToObjectID()

	mm.lock.Lock()
	mm.connections[meta.ID] = meta
	mm.lock.Unlock()

	return true
}

func (meta *ConnMeta) Send(b []byte) (int, error) {
	if meta.Conn == nil {
		return 0, errors.New("ConnMeta.Conn is nil")
	}

	n, err := meta.Conn.Write(b)
	meta.SentByte += n
	meta.SentNum += 1

	//todo
	return n, err
}

func (meta *ConnMeta) BroadCastSendClient(opName string, content []byte) {
	n, err := meta.Conn.Write(content)
	if err != nil {
		Clients.Logout(meta)
		logger.Error("brdcast: MSG <%16s> send to player %s error: %s",
			opName, meta.ID, err.Error())
		return
	}
	logger.Info("brdcast: MSG <%16s> foward to player %s success (%d bytes)",
		opName, meta.ID, n)

}

func (meta *ConnMeta) CsToClient(opName string, content []byte) bool {
	_, err := meta.Send(content)
	if err != nil {
		Servers.Logout(meta)
		logger.Error("cs: MSG <%16s> froward to player %s error: %s", opName, meta.ID, err.Error())
		return false
	}

	logger.Info("cs: MSG <%16s>: forwad to player %s sucess ", opName, meta.ID)

	return true
}

func (meta *ConnMeta) GsToServer(opName, ackName string, content []byte) bool {
	_, err := meta.Send(content)
	if err != nil {
		Servers.Logout(meta)
		logger.Error("gs: MSG <%16s> ACK <%16s> send to server %s error: %s", opName, ackName, meta.ID, err.Error())
		return false
	}

	logger.Info("gs: MSG <%16s>: server %s, ACK <%16s> send sucess", opName, meta.ID, ackName)

	return true
}

func (meta *ConnMeta) CsToServer(opName string, content []byte) bool {
	_, err := meta.Send(content)
	if err != nil {
		Servers.Logout(meta)
		logger.Error("gateway: MSG <%16s> froward to server %s error: %s", opName, meta.ID, err.Error())
		return false
	}

	logger.Info("gateway: MSG <%16s>: forwad to server %s sucess ", opName, meta.ID)

	return true
}
