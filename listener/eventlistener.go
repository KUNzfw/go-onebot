package listener

// 事件监听接口
type EventListener interface {
	Poll() (map[string]interface{}, error)
}
