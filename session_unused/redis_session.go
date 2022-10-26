package session

import (
	"encoding/json"
	"errors"
	"sync"

	"github.com/go-redis/redis"
)

//redisSession设计
//定义对象
//构造函数，为了获取对象
//Set()
//Get()
//Del()
//Save()

type RedisSession struct {
	sessionId string
	pool      *redis.Client
	//session可以先放在内存的map中
	//批量导入redis，提升性能
	sessionMap map[string]interface{}
	//读写锁
	rwlock sync.RWMutex
	//记录内存中的map是否被操作
	flag int
}

// 用常量去定义状态
const (
	SessionFlagNone = iota
	SessionFlagModify
)

func (r *RedisSession) initRedis() (err error) {
	r.pool = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	_, err = r.pool.Ping().Result()
	return
}

// 构造函数
func NewRedisSession(id string, pool *redis.Client) *RedisSession {
	s := &RedisSession{
		sessionId:  id,
		pool:       pool,
		sessionMap: make(map[string]interface{}, 16),
		flag:       SessionFlagNone,
	}
	return s
}

// Set()
func (r *RedisSession) Set(key string, value interface{}) (err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	r.sessionMap[key] = value
	//modify flag
	r.flag = SessionFlagModify
	return
}

// Save()//不知道对不对
func (r *RedisSession) Save() (err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	if r.flag != SessionFlagModify {
		return
	}
	data, err := json.Marshal(r.sessionMap)
	if err != nil {
		return
	}
	//获取redis连接
	err = r.initRedis()
	if err != nil {
		return
	}
	r.pool.Set(r.sessionId, string(data), 0)
	r.flag = SessionFlagNone
	return
}

// Get()
func (r *RedisSession) Get(key string) (value interface{}, err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	//先判断内存中有没有数据
	result, ok := r.sessionMap[key]
	if !ok {
		err = errors.New("key is not exsit")
	}
	return result,nil
}

func(r *RedisSession)Del(key string)(err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	r.flag = SessionFlagModify
	delete(r.sessionMap,key)
	return
}

func (r *RedisSession)LoadFromRedis()(err error){
	return
}
