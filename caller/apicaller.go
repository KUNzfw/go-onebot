/*
 * @Date: 2021-04-16 19:53:00
 * @LastEditors: KUNzfw
 * @LastEditTime: 2021-04-16 20:47:33
 * @FilePath: \go-onebot\caller\apicaller.go
 */
package caller

// api调用接口
type ApiCaller interface {
	Call(action string, data map[string]interface{}, result interface{}) error
}
