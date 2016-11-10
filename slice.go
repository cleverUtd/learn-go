package demo

import (
	"fmt"
)

/*
一个 slice 会指向一个序列的值，并且包含了长度信息。

[]T 是一个元素类型为 T 的 slice。

len(s) 返回 slice s 的长度。
*/


func Slice() {
    s := []int{2,3,5,7,11,13}
    fmt.Println("s ==", s)

    for i:=0;i<len(s);i++ {
        fmt.Printf("s[%d] == %d\n", i, s[i])
    }
}

//slice 由函数 make 创建。这会分配一个全是零值的数组并且返回一个 slice 指向这个数组：
/*
a := make([]int, 5)  // len(a)=5
*/
func MakeSlice() {
    a := make([]int, 5)
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)
}

/*
向 slice 添加元素
func append(s []T, vs ...T) []T
append 的第一个参数 s 是一个元素类型为 T 的 slice ，其余类型为 T 的值将会附加到该 slice 的末尾。

append 的结果是一个包含原 slice 所有元素加上新添加的元素的 slice。

如果 s 的底层数组太小，而不能容纳所有值时，会分配一个更大的数组。 返回的 slice 会指向这个新分配的数组。
*/
func AppendSlice() {
    var a []int
    printSlice("a", a)

    a = append(a, 0)
    printSlice("a", a)

    a = append(a, 2,3,4)
    printSlice("a", a)
    
}
func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
