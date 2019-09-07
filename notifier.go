/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2019/9/7
   Description :
-------------------------------------------------
*/

package zobserver

import (
    "github.com/zlyuancn/zerrors"
    "sync"
)

// 通告者储存区类型
type NotifierStorage map[string]INotifier

// 观察者储存区类型
type ObserverStorage map[IObserver]struct{}

var defaultNotifierStorage = make(NotifierStorage, InitNotifierCapacity)
var defaultSyncMutex sync.Mutex

// 获取通告者, 如果没有返回nil
func GetNotifier(name string) INotifier {
    defaultSyncMutex.Lock()
    defer defaultSyncMutex.Unlock()

    n, _ := defaultNotifierStorage[name]
    return n
}

type INotifier interface {
    // 注册观察者
    Register(IObserver)
    // 注册观察函数
    RegisterObserverFunc(fn ActionFunc) IObserver
    // 取消注册观察者
    Deregister(IObserver)
    // 通告观察者
    Notify(IMessage)
    // 通告
    NotifyMessage(body interface{})
}

type notifier struct {
    name      string
    observers ObserverStorage
    mx        sync.Mutex
}

// 注册观察者
func (m *notifier) Register(observer IObserver) {
    m.mx.Lock()
    defer m.mx.Unlock()

    m.observers[observer] = struct{}{}
}

// 注册观察函数
func (m *notifier) RegisterObserverFunc(fn ActionFunc) IObserver {
    o := NewObserver(fn)
    m.Register(o)
    return o
}

// 取消注册观察者
func (m *notifier) Deregister(observer IObserver) {
    m.mx.Lock()
    defer m.mx.Unlock()

    delete(m.observers, observer)
}

// 通告消息
func (m *notifier) Notify(msg IMessage) {
    m.mx.Lock()
    defer m.mx.Unlock()

    for o := range m.observers {
        o.OnNotify(m.name, msg)
    }
}

// 通告消息
func (m *notifier) NotifyMessage(body interface{}) {
    m.Notify(&message{
        body: body,
    })
}

// 创建一个通告者
func NewNotifier(name string) (INotifier, error) {
    n := &notifier{
        name:      name,
        observers: make(ObserverStorage, InitObserverCapacity),
    }

    defaultSyncMutex.Lock()
    defer defaultSyncMutex.Unlock()

    if _, ok := defaultNotifierStorage[name]; ok {
        return nil, zerrors.New(CreateNotifierIsExistsErr)
    }

    defaultNotifierStorage[name] = n
    return n, nil
}
