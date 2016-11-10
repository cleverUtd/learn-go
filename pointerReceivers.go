package demo


/*
方法可以与命名类型或命名类型的指针关联。

使用指针接收。有两个好处

1. 避免在每个方法调用中拷贝值（如果值类型是大的结构体的话会更有效率）
2. 方法可以修改接收者指向的值。

*/


import (
	"fmt"
	"math"
)

type VertexPR struct {
    X, Y float64
}

func (v *VertexPR) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

func (v *VertexPR) Abs() float64{
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func PointerRecv() {
    v := &VertexPR{3,4}

    fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
    v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}