package goroutine

/* sync.Mutex
只是想保证在每个时刻，只有一个 goroutine 能访问一个共享的变量从而避免冲突

可以使用 _互斥锁_(mutex)_来提供这个限制

Go 标准库中提供了 sync.Mutex 类型及其两个方法：

Lock
Unlock

我们可以通过在代码前调用 Lock 方法，在代码后调用 Unlock 方法来保证一段代码的互斥执行
也可以用 defer 语句来保证互斥锁一定会被解锁
*/
import (
	"sync"
	"time"
	"fmt"
)

type SafeCounter struct {
    v map[string]int
    mux sync.Mutex
}

//Inc 增加给定 key 的计数器的值。
func (c *SafeCounter) inc(key string) {
    c.mux.Lock()

    //Lock 之后同一时刻只有一个 goroutine 能访问 c.v
    c.v[key]++

    c.mux.Unlock()
}
func (c * SafeCounter) value(key string) int {
    c.mux.Lock()

    defer c.mux.Unlock()
	return c.v[key]
}

func SyncMutex() {
    c := SafeCounter{v: make(map[string]int)}
    for i := 0;i < 1000;i++ {
        go c.inc("somekey")
    }

    time.Sleep(time.Second)
    fmt.Println(c.value("somekey"))
}