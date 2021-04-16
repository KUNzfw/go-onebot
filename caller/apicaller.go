package caller

type ApiCaller interface {
	Call(action string, data map[string]interface{}) (map[string]interface{}, error)
}
