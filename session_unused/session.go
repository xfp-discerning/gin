package session
type Session interface{
	Set(key string ,value interface{}) error
	Del(key string) error
	Get(key string)(interface{}, error)
	Save() error 
}