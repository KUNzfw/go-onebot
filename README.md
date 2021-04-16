# go-onebot

## 快速上手
```go
// 一个简单的私聊复读机
package main

import (
	"log"

	"github.com/KUNzfw/go-onebot/bot"
)

func main() {
	wb := bot.NewWsBot("ws://localhost:6700/", nil)
	for {
		event, err := bot.PollEvent(wb)
		if err != nil {
			log.Fatal(err)
		}
		switch event.Type {
		case bot.TypePrivateMesssage:
			data := event.Data.(bot.EventPrivateMessage)
			log.Printf("Recived a private message from %v (%v): %v\n", data.Sender.Nickname, data.Sender.UserId, data.Message)
			msgId, err := bot.SendPrivateMessage(wb, data.Sender.UserId, data.Message, false)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Sent message to %v (%v): %v [Message ID: %v]", data.Sender.Nickname, data.Sender.UserId, data.Message, msgId)
		}
	}
}

```
