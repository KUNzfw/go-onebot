package onebot

import (
	"context"

	"github.com/KUNzfw/go-onebot/caller"
	"github.com/KUNzfw/go-onebot/listener"
)

// 创建WsBot的选项
type WsBotOptions struct {
	Ctx         context.Context
	AccessToken string
}

// CreateWsBot 创建一个WsBot
func CreateWsBot(url string, opts *WsBotOptions) *Bot {
	// 处理配置
	if opts == nil {
		opts = &WsBotOptions{}
	}
	if opts.Ctx == nil {
		opts.Ctx = context.Background()
	}

	// 创建caller, listener
	wsCaller := caller.CreateWsCaller(opts.Ctx, url, opts.AccessToken)
	wsListener := listener.CreateWsListener(opts.Ctx, url, opts.AccessToken)
	return &Bot{
		caller:   wsCaller,
		listener: wsListener,
	}
}
