/*
 * @Date: 2021-04-16 19:53:00
 * @LastEditors: KUNzfw
 * @LastEditTime: 2021-04-16 20:08:08
 * @FilePath: \go-onebot\bot\event.go
 */
package bot

import (
	"github.com/KUNzfw/go-onebot/listener"
	"github.com/mitchellh/mapstructure"
)

const (
	TypePrivateMesssage int32 = iota
)

// 事件类型，包含类型和数据
type Event struct {
	Type int32       // 类型
	Data interface{} // 数据
}

// PollEvent 获取事件
func PollEvent(bot listener.EventListener) (event Event, err error) {
	data, err := bot.Poll()
	event_type, event_data := parseEvent(data)
	return Event{event_type, event_data}, err
}

// parseEvent 解析事件
func parseEvent(data map[string]interface{}) (int32, interface{}) {
	if data["post_type"] == "message" {
		if data["message_type"] == "private" {
			var result EventPrivateMessage
			mapstructure.Decode(data, &result)
			return TypePrivateMesssage, result
		}
	}
	return -1, nil
}
