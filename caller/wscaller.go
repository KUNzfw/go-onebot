package caller

import (
	"context"
	"errors"
	"net/http"
	"time"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

const TIME_OUT time.Duration = time.Second * 10
const ECHO_FLAG string = "go-onebot"

type WsCaller struct {
	url          string
	access_token string
}

func CreateWsCaller(url string, access_token string) *WsCaller {
	return &WsCaller{
		url:          url,
		access_token: access_token,
	}
}

func (wc *WsCaller) Call(ctx context.Context, action string, data map[string]interface{}) (map[string]interface{}, error) {
	// 设置超时
	ctx, cancel := context.WithTimeout(ctx, TIME_OUT)
	defer cancel()

	// 处理请求头
	opts := &websocket.DialOptions{}
	opts.HTTPHeader = http.Header{}
	if wc.access_token != "" {
		opts.HTTPHeader.Add("Authorization", "Bearer "+wc.access_token)
	}

	// 建立websocket连接
	c, resp, err := websocket.Dial(ctx, wc.url, opts)

	// 检查鉴权错误
	if resp.StatusCode == 401 {
		return nil, errors.New("failed to connect: 401 unauthorized, maybe due to empty access token")
	}
	if resp.StatusCode == 403 {
		return nil, errors.New("failed to connect: maybe due to inconsistent access token")
	}
	// 其他错误
	if err != nil {
		return nil, err
	}

	defer c.Close(websocket.StatusInternalError, "internal error")

	// 编码参数，使用echo过滤生命周期回报
	wsdata := map[string]interface{}{
		"action": action,
		"params": data,
		"echo":   ECHO_FLAG,
	}

	// 发送数据
	err = wsjson.Write(ctx, c, wsdata)
	if err != nil {
		return nil, err
	}

	wsrep := make(map[string]interface{})
	// 接受回报
	for {
		err = wsjson.Read(ctx, c, &wsrep)
		if err != nil {
			return nil, err
		}
		if wsrep["echo"] == ECHO_FLAG {
			break
		}
	}

	c.Close(websocket.StatusNormalClosure, "")

	// 检测400和404错误
	if retcode := wsrep["retcode"].(float64); retcode == 1404.0 {
		return nil, errors.New("failed to call: 404 not found")
	} else if retcode == 1400.0 {
		return nil, errors.New("failed to call: 400 bad request")
	}

	// 提取数据并进行类型转换
	if rep, ok := wsrep["data"].(map[string]interface{}); ok {
		return rep, nil
	} else {
		return nil, nil
	}
}
