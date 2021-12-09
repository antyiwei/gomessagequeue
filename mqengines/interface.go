package mqengines

//对于中间件引擎的抽象，随着后续模式的增加，做灵活调整

type MqEngine interface {
	Add(key string, data interface{}) error
	Read(key string, count int64, onRead func(ds []interface{}) error)
}
