package session

import (
	"sync"

	uuid "github.com/satori/go.uuid"
)

//定义MemorySessionMgr对象（存放所有session的map，读写锁）
//构造函数
//Init()
//CreateSession()
//GetSession()

type MemorySessionMgr struct {
	sessionMap map[string]Session
	rwlock     sync.RWMutex
}

func NewMemorySessionMgr() *MemorySessionMgr {
	sr := &MemorySessionMgr{
		sessionMap: make(map[string]Session, 1024),
	}
	return sr
}

func (m *MemorySessionMgr) Init(addr string, options ...string) (err error) {
	return
}

func (m *MemorySessionMgr) CreateSession() (session Session, err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	//使用github中的go.uuid包
	id := uuid.NewV4()
	//将id转成string
	sessionId := id.String()
	ms := NewMemorySession(sessionId)
}
