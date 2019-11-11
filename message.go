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
    Type() string
    // 获取通告主要内容
    Body() interface{}
    // 获取通告附加数据
    Meta() interface{}
}

type message struct {
    msg_type string
    body     interface{}
    meta     interface{}
}

// 获取通告消息类型
func (m *message) Type() string {
    return m.msg_type
}

// 获取通告主要内容
func (m *message) Body() interface{} {
    return m.body
}

// 获取通告自定义数据
func (m *message) Meta() interface{} {
    return m.meta
}

// 创建一条消息
func NewMessage(body interface{}) IMessage {
    return &message{
        body: body,
    }
}

// 创建一条指定类型的消息
func NewMessageWithType(msg_type string, body interface{}) IMessage {
    return &message{
        msg_type: msg_type,
        body:     body,
    }
}

// 创建一条包含附加数据的消息
func NewMessageWithMeta(msg_type string, body interface{}, meta interface{}) IMessage {
    return &message{
        msg_type: msg_type,
        body:     body,
        meta:     meta,
    }
}
