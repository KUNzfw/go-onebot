/*
 * @Date: 2021-04-16 19:53:00
 * @LastEditors: KUNzfw
 * @LastEditTime: 2021-04-16 20:08:34
 * @FilePath: \go-onebot\bot\eventtype.go
 */
package bot

// 发送者
type Sender struct {
	UserId   int64  `mapstructure:"user_id"`
	Nickname string `mapstructure:"nickname"`
	Sex      string `mapstructure:"sex"`
	Age      int32  `mapstructure:"age"`
}

// 私聊消息类型
type EventPrivateMessage struct {
	Time        int64  `mapstructure:"time"`
	SelfID      int64  `mapstructure:"self_id"`
	PostType    string `mapstructure:"post_type"`
	MessageType string `mapstructure:"message_type"`
	SubType     string `mapstructure:"sub_type"`
	MessageID   int32  `mapstructure:"message_id"`
	UserId      int64  `mapstructure:"user_id"`
	Message     string `mapstructure:"message"`
	RawMessage  string `mapstructure:"raw_message"`
	Font        int32  `mapstructure:"font"`
	Sender      Sender `mapstructure:"sender"`
}
