package demo

import (
	"fmt"
)

func ForLoop() {
    sum := 0
    for i:=0;i<10;i++ {
        sum++
    }
    fmt.Println(sum)
}

//ForLoop1 循环初始化语句和后置语句都是可选的。
func ForLoop1() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}

//WhileLoop C 的 while 在 Go 中叫做 for 。
func WhileLoop()  {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}