/*
 * @Date: 2021-04-16 19:53:00
 * @LastEditors: KUNzfw
 * @LastEditTime: 2021-04-16 20:14:55
 * @FilePath: \go-onebot\listener\wslistener.go
 */
package listener

import (
	"context"
	"errors"
	"net/http"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

// 通过websocket进行事件监听的封装
type WsListener struct {
	ctx          context.Context
	url          string
	access_token string
	rec_chan     <-chan map[string]interface{}
	send_chan    chan<- map[string]interface{}
	err_chan     chan error
	is_serve     bool
}

// CreateWsListener 创建WsListener实例
func CreateWsListener(url string, access_token string, ctx context.Context) *WsListener {
	wschan := make(chan map[string]interface{}, 64)
	err_chan := make(chan error, 8)
	return &WsListener{
		ctx:          ctx,
		url:          url,
		access_token: access_token,
		rec_chan:     wschan,
		send_chan:    wschan,
		err_chan:     err_chan,
		is_serve:     false,
	}
}

// Poll 实现Poll接口
func (wl *WsListener) Poll() (map[string]interface{}, error) {
	// 启动事件监听服务
	if !wl.is_serve {
		go wl.serve()
	}

	select {
	case data := <-wl.rec_chan:
		return data, nil
	case <-wl.ctx.Done():
		return nil, nil
	case err := <-wl.err_chan:
		return nil, err
	}
}

func (wl *WsListener) serve() {
	// 管理服务状态
	wl.is_serve = true

	// 处理请求头
	opts := &websocket.DialOptions{}
	opts.HTTPHeader = http.Header{}
	if wl.access_token != "" {
		opts.HTTPHeader.Add("Authorization", "Bearer "+wl.access_token)
	}

	// 建立websocket连接
	c, resp, err := websocket.Dial(wl.ctx, wl.url, opts)

	// 检查鉴权错误
	if resp.StatusCode == 401 {
		wl.err_chan <- errors.New("failed to connect: 401 unauthorized, maybe due to empty access token")
		return
	}
	if resp.StatusCode == 403 {
		wl.err_chan <- errors.New("failed to connect: 403 forbidden, maybe due to inconsistent access token")
		return
	}
	// 其他错误
	if err != nil {
		wl.err_chan <- err
		return
	}

	defer c.Close(websocket.StatusInternalError, "internal error")
	defer func() {
		wl.is_serve = false
	}()

	// 读取数据并发送到chan
	for {
		select {
		case <-wl.ctx.Done():
			c.Close(websocket.StatusNormalClosure, "")
		default:
			data := make(map[string]interface{})
			err := wsjson.Read(wl.ctx, c, &data)
			if err != nil {
				wl.err_chan <- err
				return
			}
			wl.send_chan <- data
		}
	}
}
