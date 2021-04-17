package caller

// api调用接口
type APICaller interface {
	Call(action string, data map[string]interface{}, result interface{}) error
}
