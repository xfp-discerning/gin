package session
//定义管理者，管理所以session
type SessionMgr interface{
	Init(addr string, options ...string)(err error)
	CreateSession()(session Session,err error)
	GetSession(sessionId string)(session Session, err error)
}