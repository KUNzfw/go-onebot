<!--
 * @Date: 2021-04-10 18:10:35
 * @LastEditors: KUNzfw
 * @LastEditTime: 2021-04-17 09:57:18
 * @FilePath: \go-onebot\README.md
-->
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
	handler := bot.EventHandler{
		OnPrivateMessage: func(data *bot.EventPrivateMessage) {
			log.Printf("收到来自 %v (%v) 的私聊消息: %v\n", data.Sender.Nickname, data.Sender.UserId, data.Message)
			msgId, err := bot.SendPrivateMessage(wb, data.Sender.UserId, data.Message, false)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("给 %v (%v) 发送私聊消息: %v [Message ID: %v]", data.Sender.Nickname, data.Sender.UserId, data.Message, msgId)
		},
	}
	if err := handler.HandleEvent(wb); err != nil {
		log.Fatal(err)
	}
}

```
