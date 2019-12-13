/*
-------------------------------------------------
   Author :       zlyuan
   date：         2019/9/7
   Description :
-------------------------------------------------
*/

package zobserver

type ActionFunc func(notifyName string, msg IMessage)

type IObserver interface {
    // 触发事件
    OnNotify(notifyName string, msg IMessage)
}

type observer struct {
    action      ActionFunc
    listen_type string
    has_type    bool
}

func (m *observer) OnNotify(notifyName string, msg IMessage) {
    if m.has_type || m.listen_type == msg.Type() {
        m.action(notifyName, msg)
    }
}

// 创建一个观察者
func NewObserver(fn ActionFunc) IObserver {
    return &observer{
        action: fn,
    }
}

// 创建一个观察者并指定监听的消息类型
func NewObserverWithType(msg_type string, fn ActionFunc) IObserver {
    return &observer{
        action:      fn,
        listen_type: msg_type,
        has_type:    true,
    }
}

// 创建一个观察者并注册到通告者
func NewObserverAndReg(notifyName string, fn ActionFunc) (INotifier, IObserver) {
    ob := &observer{
        action: fn,
    }
    notifier := CreateOrGerNotifier(notifyName)
    notifier.Register(ob)
    return notifier, ob
}

// 创建一个观察者并指定监听的消息类型然后注册到通告者
func NewObserverAndRegWithType(notifyName string, msg_type string, fn ActionFunc) (INotifier, IObserver) {
    ob := &observer{
        action:      fn,
        listen_type: msg_type,
        has_type:    true,
    }
    notifier := CreateOrGerNotifier(notifyName)
    notifier.Register(ob)
    return notifier, ob
}
