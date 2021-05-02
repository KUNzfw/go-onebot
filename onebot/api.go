package onebot

import "github.com/mitchellh/mapstructure"

// SendPrivateMessage 发送私聊消息
func (bot *Bot) SendPrivateMessage(userID int64, message string, autoEscape bool) (messageID int32, err error) {
	resp := make(map[string]interface{})
	if cerr := bot.caller.Call("send_private_msg", map[string]interface{}{
		"user_id":     userID,
		"message":     message,
		"auto_escape": autoEscape,
	}, &resp); cerr != nil {
		return 0, cerr
	}

	err = nil
	if id, ok := resp["message_id"].(float64); ok {
		// 根据onebot标准，这里的强制转换没有问题
		messageID = int32(id)
	}
	return
}

type LoginInfo struct {
	UserID   int64  `mapstructure:"user_id"`
	NickName string `mapstructure:"nickname"`
}

// GetLoginInfo 获取登录号信息
func (bot *Bot) GetLoginInfo() (info *LoginInfo, err error) {
	resp := make(map[string]interface{})
	if cerr := bot.caller.Call("get_login_info", nil, &resp); cerr != nil {
		return nil, cerr
	}

	info = &LoginInfo{}
	err = mapstructure.Decode(resp, info)
	return
}
