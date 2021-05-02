package listener

import (
	"context"
	"errors"
	"net/http"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

const (
	wsChanCap  int = 64
	errChanCap int = 8
)

// 通过websocket进行事件监听的封装
type WsListener struct {
	ctx         context.Context
	url         string
	accessToken string
	recChan     <-chan map[string]interface{}
	sendChan    chan<- map[string]interface{}
	errChan     chan error
	isServe     bool
}

// CreateWsListener 创建WsListener实例
func CreateWsListener(ctx context.Context, url, accessToken string) *WsListener {
	wschan := make(chan map[string]interface{}, wsChanCap)
	errchan := make(chan error, errChanCap)
	return &WsListener{
		ctx:         ctx,
		url:         url,
		accessToken: accessToken,
		recChan:     wschan,
		sendChan:    wschan,
		errChan:     errchan,
		isServe:     false,
	}
}

// Poll 实现Poll接口
func (wl *WsListener) Poll() (map[string]interface{}, error) {
	// 启动事件监听服务
	if !wl.isServe {
		go wl.serve()
	}

	select {
	case data := <-wl.recChan:
		return data, nil
	case <-wl.ctx.Done():
		return nil, nil
	case err := <-wl.errChan:
		return nil, err
	}
}

func (wl *WsListener) serve() {
	// 管理服务状态
	wl.isServe = true

	// 处理请求头
	opts := &websocket.DialOptions{}
	opts.HTTPHeader = http.Header{}
	if wl.accessToken != "" {
		opts.HTTPHeader.Add("Authorization", "Bearer "+wl.accessToken)
	}

	// 建立websocket连接
	c, resp, err := websocket.Dial(wl.ctx, wl.url, opts)

	if err != nil {
		if resp != nil {
			// 检查鉴权错误
			if resp.StatusCode == http.StatusUnauthorized {
				wl.errChan <- errors.New("事件服务器连接失败: 401 Unauthorized, 可能因为访问密钥未提供")
				return
			}
			if resp.StatusCode == http.StatusForbidden {
				wl.errChan <- errors.New("事件服务器连接失败: 403 Forbidden, 可能因为访问密钥错误")
				return
			}
		}

		// 其他错误
		wl.errChan <- errors.New("事件服务器连接失败: " + err.Error())
		return
	}

	defer c.Close(websocket.StatusInternalError, "internal error")
	defer func() {
		wl.isServe = false
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
				wl.errChan <- errors.New("事件读取错误: " + err.Error())
				return
			}
			wl.sendChan <- data
		}
	}
}
