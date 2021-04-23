package bot

import (
	"github.com/KUNzfw/go-onebot/caller"
	"github.com/KUNzfw/go-onebot/listener"
)

// 用于方便地调用各种api和接受各种event
type Bot struct {
	caller   caller.APICaller
	listener listener.EventListener
}
