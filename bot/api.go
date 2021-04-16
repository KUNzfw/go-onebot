/*
 * @Date: 2021-04-16 19:53:00
 * @LastEditors: KUNzfw
 * @LastEditTime: 2021-04-16 20:07:43
 * @FilePath: \go-onebot\bot\api.go
 */
package bot

import (
	"github.com/KUNzfw/go-onebot/caller"
)

// SendPrivateMessage 发送私聊消息
func SendPrivateMessage(bot caller.ApiCaller, userId int64, message string, auto_escape bool) (messageId int32, err error) {
	rep, cerr := bot.Call("send_private_msg", map[string]interface{}{
		"user_id":     userId,
		"message":     message,
		"auto_escape": auto_escape,
	})

	if cerr != nil {
		return 0, cerr
	}

	err = nil
	if id, ok := rep["message_id"].(float64); ok {
		// 根据onebot标准，这里的强制转换没有问题
		messageId = int32(id)
	}
	return
}
