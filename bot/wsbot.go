package bot

import (
	"context"

	"github.com/KUNzfw/go-onebot/caller"
	"github.com/KUNzfw/go-onebot/listener"
)

// 创建WsBot的选项
type WsBotOptions struct {
	ctx         context.Context
	accessToken string
}

// NewWsBot 创建一个WsBot
func NewWsBot(url string, opts *WsBotOptions) *Bot {
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
	return &Bot{
		caller:   wsCaller,
		listener: wsListener,
	}
}
