/*
-------------------------------------------------
   Author :       zlyuan
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
    // 销毁
    Distroy()
}

type notifier struct {
    name      string
    observers ObserverStorage
    distroy   bool
    mx        sync.Mutex
}

// 注册观察者
func (m *notifier) Register(observer IObserver) {
    m.mx.Lock()
    defer m.mx.Unlock()

    if !m.distroy {
        m.observers[observer] = struct{}{}
    }
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

    if !m.distroy {
        for o := range m.observers {
            o.OnNotify(m.name, msg)
        }
    }
}

// 通告消息
func (m *notifier) NotifyMessage(body interface{}) {
    m.Notify(&message{
        body: body,
    })
}

// 销毁
func (m *notifier) Distroy() {
    m.mx.Lock()
    defer m.mx.Unlock()

    m.distroy = true
    m.observers = ObserverStorage{}
}

// 创建一个通告者
func NewNotifier(name string) (INotifier, error) {
    defaultSyncMutex.Lock()
    defer defaultSyncMutex.Unlock()

    if _, ok := defaultNotifierStorage[name]; ok {
        return nil, zerrors.New(CreateNotifierIsExistsErr)
    }

    n := &notifier{
        name:      name,
        observers: make(ObserverStorage, InitObserverCapacity),
    }
    defaultNotifierStorage[name] = n
    return n, nil
}

// 创建一个通告者, 如果通告者已存在则获取它
func CreateOrGerNotifier(name string) INotifier {
    defaultSyncMutex.Lock()
    defer defaultSyncMutex.Unlock()

    if notifier, ok := defaultNotifierStorage[name]; ok {
        return notifier
    }

    n := &notifier{
        name:      name,
        observers: make(ObserverStorage, InitObserverCapacity),
    }
    defaultNotifierStorage[name] = n
    return n
}

// 获取通告者, 如果没有返回nil
func GetNotifier(name string) INotifier {
    defaultSyncMutex.Lock()
    defer defaultSyncMutex.Unlock()

    n, _ := defaultNotifierStorage[name]
    return n
}

// 删除通告者
func DelNotifier(name string) {
    defaultSyncMutex.Lock()
    defer defaultSyncMutex.Unlock()

    if notifier, ok := defaultNotifierStorage[name]; ok {
        delete(defaultNotifierStorage, name)
        notifier.Distroy()
    }
}
