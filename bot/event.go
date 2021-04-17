package bot

import (
	"errors"

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

// 事件处理器
type EventHandler struct {
	OnPrivateMessage func(data *EventPrivateMessage)
}

// HandleEvent 监听并处理事件
func (handler *EventHandler) HandleEvent(bot listener.EventListener) error {
	for {
		rawdata, err := bot.Poll()
		if err != nil {
			return errors.New("监听事件时发生错误: " + err.Error())
		}

		switch parseEvent(rawdata) {
		case TypePrivateMesssage:
			data := &EventPrivateMessage{}
			if err := mapstructure.Decode(rawdata, data); err == nil {
				handler.OnPrivateMessage(data)
			}
		default:
		}
	}
}

// parseEvent 解析事件
func parseEvent(data map[string]interface{}) int32 {
	if data["post_type"] == "message" {
		if data["message_type"] == "private" {
			return TypePrivateMesssage
		}
	}
	return -1
}
