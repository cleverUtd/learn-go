package demo

import (
	"fmt"
)

/*
变量在定义时没有明确的初始化时会赋值为 零值 。

数值类型为 0 ，
布尔类型为 false ，
字符串为 "" （空字符串）。
*/


//var 语句可以定义在包或函数级别
var c, python, java bool
var i, j int = 1, 2

//VariableDeclairation 变量定义
func VariableDeclairation()  {
    var i int
    fmt.Println(i, c, python, java)
}

//InitVar 变量定义可以包含初始值，每个变量对应一个。如果初始化是使用表达式，则可以省略类型；变量从初始值中获得类型。
func InitVar() {
    var c, python, java = true, false, "no!"
    fmt.Println(i, j, c, python, java)
}

//shortInit  := 简洁赋值语句在明确类型的地方，可以用于替代 var 定义
/*
函数外的每个语句都必须以关键字开始（ var 、 func 、等等）， := 结构不能使用在函数外。
*/
func shortInit() {
    var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

    fmt.Println(i, j, k, c, python, java)
}