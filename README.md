# go-onebot

## 快速上手
```go
// 一个简单的私聊复读机
package main

import (
	"context"
	"log"

	"github.com/KUNzfw/go-onebot/bot"
)

func main() {
	wb := bot.NewWsBot(context.Background(), "ws://localhost:6700/", "")
	for {
		event_type, event_data, err := bot.PollEvent(wb)
		if err != nil {
			log.Fatal(err)
		}
		switch event_type {
		case bot.TypePrivateMesssage:
			data := event_data.(bot.EventPrivateMessage)
			log.Printf("Recived a private message from %v (%v): %v\n", data.Sender.Nickname, data.Sender.UserId, data.Message)
			msgId, err := bot.SendPrivateMessage(wb, context.Background(), data.Sender.UserId, data.Message, false)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Sent message to %v (%v): %v \tMessage ID: %v", data.Sender.Nickname, data.Sender.UserId, data.Message, msgId)
		}
	}
}
```
