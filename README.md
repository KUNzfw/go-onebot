# go-onebot

## 快速上手
```go
// 一个简单的私聊复读机
package main

import (
	"log"

	"github.com/KUNzfw/go-onebot/onebot"
)

func main() {
	bot := onebot.CreateWsBot("ws://localhost:6700/", nil)
	handler := onebot.EventHandler{
		OnPrivateMessage: func(data *onebot.EventPrivateMessage) {
			log.Printf("收到来自 %v (%v) 的私聊消息: %v\n", data.Sender.Nickname, data.Sender.UserID, data.Message)
			msgID, err := bot.SendPrivateMessage(data.Sender.UserID, data.Message, false)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("给 %v (%v) 发送私聊消息: %v [Message ID: %v]", data.Sender.Nickname, data.Sender.UserID, data.Message, msgID)
		},
	}
	if err := bot.HandleEvent(&handler); err != nil {
		log.Fatal(err)
	}
}
```
