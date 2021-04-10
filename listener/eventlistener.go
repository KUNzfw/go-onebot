package listener

type EventListener interface {
	Poll() (map[string]interface{}, error)
}
