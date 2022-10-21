package session

import "sync"

//MemorySessio的设计
//定义MemorySession对象（字段：sessionId，存k-v的map，读写锁）
//构造函数为了获取对象
//Set()
//Get()
//Del()
//Save()

type MemorySession struct{
	sessionId string
	data map[string]interface{}
	rwlock sync.RWMutex
}

//构造函数
 func NewMemorySession(id string) *MemorySession{
	s := &MemorySession{
		sessionId: id,
		data: make(map[string]interface{},16),
	}
	return s
 }