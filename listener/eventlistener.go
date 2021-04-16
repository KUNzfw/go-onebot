/*
 * @Date: 2021-04-16 19:53:00
 * @LastEditors: KUNzfw
 * @LastEditTime: 2021-04-16 20:12:45
 * @FilePath: \go-onebot\listener\eventlistener.go
 */
package listener

// 事件监听接口
type EventListener interface {
	Poll() (map[string]interface{}, error)
}
