package demo

/* 
首字母大写的名称是被导出的, 在导入包之后，你只能访问包所导出的名字，任何未导出的名字是不能被包外的代码访问的 
例如Sum和SUM都是被导出的名称，sum不会被导出
*/

//Sum is used to calculate the sum of two integer numbers
func Sum(a int, b int) int{
    return a+b;
}

//Add 当两个或多个连续的函数命名参数是同一类型，则除了最后一个类型之外，其他都可以省略。
func Add(x, y int) int {
    return x+y
}

//Swap 函数返回了两个字符串. 函数可以返回任意数量的返回值。
func Swap(x, y string) (string, string) {
    return y,x
}

