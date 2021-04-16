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

type WsBotOptions struct {
	ctx          context.Context
	access_token string
}

func NewWsBot(url string, opts *WsBotOptions) *WsBot {
	// 处理配置
	if opts == nil {
		opts = &WsBotOptions{}
	}
	if opts.ctx == nil {
		opts.ctx = context.Background()
	}

	// 创建caller, listener
	caller := caller.CreateWsCaller(url, opts.access_token, opts.ctx)
	listener := listener.CreateWsListener(url, opts.access_token, opts.ctx)
	return &WsBot{
		caller:   caller,
		listener: listener,
	}
}

func (bot *WsBot) Call(action string, data map[string]interface{}) (map[string]interface{}, error) {
	return bot.caller.Call(action, data)
}

func (bot *WsBot) Poll() (map[string]interface{}, error) {
	return bot.listener.Poll()
}
