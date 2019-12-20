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
    action     ActionFunc
    msgtypes   map[string]struct{}
    is_msgtype bool
}

func (m *observer) OnNotify(notifyName string, msg IMessage) {
    if !m.is_msgtype {
        m.action(notifyName, msg)
        return
    }

    if _, ok := m.msgtypes[msg.Type()]; ok {
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
func NewObserverWithType(fn ActionFunc, msgtypes ...string) IObserver {
    mm := make(map[string]struct{}, len(msgtypes))
    for _, t := range msgtypes {
        mm[t] = struct{}{}
    }
    return &observer{
        action:     fn,
        msgtypes:   mm,
        is_msgtype: true,
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
func NewObserverAndRegWithType(notifyName string, fn ActionFunc, msgtypes ...string) (INotifier, IObserver) {
    mm := make(map[string]struct{}, len(msgtypes))
    for _, t := range msgtypes {
        mm[t] = struct{}{}
    }
    ob := &observer{
        action:     fn,
        msgtypes:   mm,
        is_msgtype: true,
    }
    notifier := CreateOrGerNotifier(notifyName)
    notifier.Register(ob)
    return notifier, ob
}
