/*
-------------------------------------------------
   Author :       Zhang Fan
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
    action ActionFunc
}

func (m *observer) OnNotify(notifyName string, msg IMessage) {
    m.action(notifyName, msg)
}

// 创建一个观察者
func NewObserver(fn ActionFunc) IObserver {
    return &observer{
        action: fn,
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
