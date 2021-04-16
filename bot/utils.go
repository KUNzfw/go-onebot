/*
 * @Date: 2021-04-16 19:53:00
 * @LastEditors: KUNzfw
 * @LastEditTime: 2021-04-16 20:10:26
 * @FilePath: \go-onebot\bot\utils.go
 */
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
	ctx          context.Context
	access_token string
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
	caller := caller.CreateWsCaller(url, opts.access_token, opts.ctx)
	listener := listener.CreateWsListener(url, opts.access_token, opts.ctx)
	return &WsBot{
		caller:   caller,
		listener: listener,
	}
}

// Call 实现Call接口
func (bot *WsBot) Call(action string, data map[string]interface{}) (map[string]interface{}, error) {
	return bot.caller.Call(action, data)
}

// Poll 实现Poll接口
func (bot *WsBot) Poll() (map[string]interface{}, error) {
	return bot.listener.Poll()
}
