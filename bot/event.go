package bot

import (
	"github.com/KUNzfw/go-onebot/listener"
	"github.com/mitchellh/mapstructure"
)

const (
	TypePrivateMesssage int32 = iota
)

func PollEvent(bot listener.EventListener) (event_type int32, event_data interface{}, err error) {
	data, err := bot.Poll()
	event_type, event_data = parseEvent(data)
	return event_type, event_data, err
}

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
