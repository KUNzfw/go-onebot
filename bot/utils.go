package bot

import (
	"context"

	"github.com/KUNzfw/go-onebot/caller"
	"github.com/KUNzfw/go-onebot/listener"
)

// 提供对WsCaller和WsListener的封装
type WsBot struct {
	caller   *caller.WsCaller
	listener *listener.WsListener
}

// 创建WsBot的选项
type WsBotOptions struct {
	ctx         context.Context
	accessToken string
}

// NewWsBot 创建一个WsBot
func NewWsBot(url string, opts *WsBotOptions) *WsBot {
	// 处理配置
	if opts == nil {
		opts = &WsBotOptions{}
	}
	if opts.ctx == nil {
		opts.ctx = context.Background()
	}

	// 创建caller, listener
	wsCaller := caller.CreateWsCaller(opts.ctx, url, opts.accessToken)
	wsListener := listener.CreateWsListener(opts.ctx, url, opts.accessToken)
	return &WsBot{
		caller:   wsCaller,
		listener: wsListener,
	}
}

// Call 实现Call接口
func (bot *WsBot) Call(action string, data map[string]interface{}, result interface{}) error {
	return bot.caller.Call(action, data, result)
}

// Poll 实现Poll接口
func (bot *WsBot) Poll() (map[string]interface{}, error) {
	return bot.listener.Poll()
}
