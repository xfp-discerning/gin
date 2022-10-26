package session

import (
	"errors"
	"sync"
)

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

 func (m *MemorySession)Set(key string, value interface{})(err error){
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	m.data[key]= value
	return
 }

 func (m *MemorySession)Get(key string)(interface{}, error){
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	value,ok := m.data[key]
	if !ok{
		err := errors.New("key is not exist in map")
		return nil,err
	}
	return value,nil
 }

 func (m *MemorySession)Del(key string)(err error){
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	delete(m.data,key)
	return
 }
 func (m *MemorySession)Save()(err error){
	return
 }