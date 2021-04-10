package bot

import (
	"context"

	"github.com/KUNzfw/go-onebot/caller"
	"github.com/KUNzfw/go-onebot/listener"
)

type WsBot struct {
	caller   *caller.WsCaller
	listener *listener.WsListener
}

func NewWsBot(ctx context.Context, url string, access_token string) *WsBot {
	caller := caller.CreateWsCaller(url, access_token)
	listener := listener.CreateWsListener(ctx, url, access_token)
	return &WsBot{
		caller:   caller,
		listener: listener,
	}
}

func (bot *WsBot) Call(ctx context.Context, action string, data map[string]interface{}) (map[string]interface{}, error) {
	return bot.caller.Call(ctx, action, data)
}

func (bot *WsBot) Poll() (map[string]interface{}, error) {
	return bot.listener.Poll()
}
