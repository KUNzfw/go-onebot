package onebot

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
