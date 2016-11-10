package goroutine

/*
channel 是有类型的管道，可以用 channel 操作符 <- 对其发送或者接收值。

ch <- v    // 将 v 送入 channel ch。
v := <-ch  // 从 ch 接收，并且赋值给 v。

（“箭头”就是数据流的方向。）

和 map 与 slice 一样，channel 使用前必须创建：

ch := make(chan int)

默认情况下，在另一端准备好之前，发送和接收都会阻塞。这使得 goroutine 可以在没有明确的锁或竞态变量的情况下进行同步。

*/

import "fmt"

func sum(a []int, c chan int)  {
    sum := 0
    for _, v := range a {
        sum += v
    }
    c <- sum //将和送入 管道c
}

func Channel() {
    a := []int{7,2,8,-9,4,0}

    c := make(chan int)
    go sum(a[:len(a) / 2], c)
    go sum(a[len(a)/2:], c)

    x, y := <-c, <-c // 从 c 中获取

    fmt.Println(x, y, x+y)
}

/* range 和 close

发送者可以 close 一个 channel 来表示再没有值会被发送了

接收者可以通过赋值语句的第二参数来测试 channel 是否被关闭：当没有值可以接收并且 channel 已经被关闭
v, ok := <-ch  ok 会被设置为 false。
*/
func fibonacci(n int, c chan int) {
    x, y := 0, 1
    for i := 0; i < n; i++ {
        c <- x
        x, y = y, x+y
    }
    close(c)
}

func Calfibonacci() {
    c := make(chan int, 10)
    go fibonacci(cap(c), c)
    // 循环 `for i := range c` 会不断从 channel 接收值，直到它被关闭。
    for i := range c {
		fmt.Println(i)
	}
}