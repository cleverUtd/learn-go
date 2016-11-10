package goroutine

import (
	"time"
	"fmt"
)


/*
goroutine 是由 Go 运行时环境管理的轻量级线程。

go f(x, y, z)

goroutine 在相同的地址空间中运行，因此访问共享内存必须进行同步
*/

func say(s string) {
    for i :=0;i<5;i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}

func GoSay() {
    go say("world")
    say("hello")
}
