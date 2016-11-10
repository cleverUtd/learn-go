package demo

import (
	"fmt"
)

/*
类型 [n]T 是一个有 n 个类型为 T 的值的数组。

表达式

var a [10]int 定义变量 a 是一个有十个整数的数组。
*/
func Array() {
    var a [2]string
    a[0] = "Hello"
    a[1] = "world"

    fmt.Println(a)
}