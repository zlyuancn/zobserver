/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2019/9/7
   Description :
-------------------------------------------------
*/

package zobserver

// 通告消息
type IMessage interface {
    // 获取通告消息类型
    MsgType() int
    // 获取通告主要内容
    Body() interface{}
    // 获取通告自定义数据
    Custom() map[string]interface{}
}

type message struct {
    msg_type int
    body     interface{}
    custom   map[string]interface{}
}

// 获取通告消息类型
func (m *message) MsgType() int {
    return m.msg_type
}

// 获取通告主要内容
func (m *message) Body() interface{} {
    return m.body
}

// 获取通告自定义数据
func (m *message) Custom() map[string]interface{} {
    return m.custom
}

// 创建一条消息
func NewMessage(body interface{}) IMessage {
    return &message{
        body: body,
    }
}

// 创建一条指定类型的消息
func NewMessageWithType(msg_type int, body interface{}) IMessage {
    return &message{
        msg_type: msg_type,
        body:     body,
    }
}

// 创建一条自定义消息
func NewCustomMessage(msg_type int, body interface{}, custom map[string]interface{}) IMessage {
    return &message{
        msg_type: msg_type,
        body:     body,
        custom:   custom,
    }
}
