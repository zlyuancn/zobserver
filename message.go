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
    // 获取通告内容
    Body() interface{}
}

type message struct {
    body interface{} //通告内容
}

// 获取通告数据
func (m *message) Body() interface{} {
    return m.body
}

// 创建一条消息
func NewMessage(body interface{}) IMessage {
    return &message{
        body: body,
    }
}
