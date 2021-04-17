package caller

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/mitchellh/mapstructure"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

const (
	TimeOut       time.Duration = time.Second * 10
	EchoFlag      string        = "go-onebot"
	RetCodeOffset int           = 1000
)

// 通过websocket进行api调用的封装
type WsCaller struct {
	url         string
	accessToken string
	ctx         context.Context
}

// api调用的返回数据
type responseData struct {
	Status  string
	Retcode int
	Data    interface{}
	Echo    string
}

// CreateWsCaller 创建WsCaller实例
func CreateWsCaller(ctx context.Context, url, accessToken string) *WsCaller {
	return &WsCaller{
		url:         url,
		accessToken: accessToken,
		ctx:         ctx,
	}
}

// Call 实现Call接口
func (wc *WsCaller) Call(action string, data map[string]interface{}, result interface{}) error {
	// 设置超时
	ctx, cancel := context.WithTimeout(wc.ctx, TimeOut)
	defer cancel()

	// 处理请求头
	opts := &websocket.DialOptions{}
	opts.HTTPHeader = http.Header{}
	if wc.accessToken != "" {
		opts.HTTPHeader.Add("Authorization", "Bearer "+wc.accessToken)
	}

	// 建立websocket连接
	c, resp, err := websocket.Dial(ctx, wc.url, opts)

	// 检查鉴权错误
	if resp.StatusCode == http.StatusUnauthorized {
		return errors.New("API服务器连接失败: 401 Unauthorized, 可能因为访问密钥未提供")
	}
	if resp.StatusCode == http.StatusForbidden {
		return errors.New("API服务器连接失败: 403 Forbidden, 可能因为访问密钥错误")
	}
	// 其他错误
	if err != nil {
		return errors.New("API服务器连接失败: " + err.Error())
	}

	defer c.Close(websocket.StatusInternalError, "internal error")

	// 编码参数，使用echo过滤生命周期回报
	wsdata := map[string]interface{}{
		"action": action,
		"params": data,
		"echo":   EchoFlag,
	}

	// 发送数据
	err = wsjson.Write(ctx, c, wsdata)
	if err != nil {
		return errors.New("向API服务器发送数据失败: " + err.Error())
	}

	rawResult := make(map[string]interface{})
	// 接受回报
	for {
		err = wsjson.Read(ctx, c, &rawResult)
		if err != nil {
			return errors.New("从服务器读取数据失败: " + err.Error())
		}
		if rawResult["echo"] == EchoFlag {
			break
		}
	}

	c.Close(websocket.StatusNormalClosure, "")

	// 将数据转换为结构体
	respResult := responseData{
		Data: result,
	}
	if err := mapstructure.Decode(rawResult, &respResult); err != nil {
		return errors.New("解析API调用返回失败: " + err.Error())
	}

	// 检测400和404错误
	switch respResult.Retcode {
	case http.StatusNotFound + RetCodeOffset:
		return errors.New("API调用失败: 404 Not Found")
	case http.StatusBadRequest + RetCodeOffset:
		return errors.New("API调用失败: 400 Bad Request")
	}

	return nil
}
