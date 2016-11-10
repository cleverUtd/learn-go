package demo

/*
接口类型是由一组方法定义的集合。

接口类型的值可以存放实现这些方法的任何值。
*/

import (
	"math"
	"fmt"
)

type Abser interface {
    Abs() float64
}

type MyFloatI float64

func (f MyFloatI) Abs() float64 {
    if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type VertexI struct {
    X, Y float64
}


func (v *VertexI) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func ImplOfInterface() {
    var a Abser

    f := MyFloatI(-math.Sqrt2)
    v := VertexI{3,4}

    a = f
    fmt.Println(a.Abs())
    a = &v
    fmt.Println(a.Abs())
}