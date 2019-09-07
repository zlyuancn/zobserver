# zobserver
###### 观察者模块

## 获得zobserver
`go get -u github.com/zlyuancn/zobserver`

## 使用zobserver

```go
package main

import (
    "fmt"
    "github.com/zlyuancn/zobserver"
)

func main() {
    // 创建一个通告者
    n, _ := zobserver.NewNotifier("notifier")

    // 创建一个观察者
    o := zobserver.NewObserver(func(notifyName string, msg zobserver.IMessage) {
        fmt.Println(notifyName, msg.Body())
    })

    // 注册观察者
    n.Register(o)

    // 通告信息
    n.NotifyMessage("hello")
}
```
