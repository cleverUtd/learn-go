package demo

/*
方法

Go 没有类。可以在结构体类型上定义方法。

方法接收者 出现在 func 关键字和方法名之间的参数中。
*/
import (
	"math"
	"fmt"
)


type Ver struct {
    A, B float64
}

func (v *Ver) Abs() float64 {
    return math.Sqrt(v.A*v.A + v.B * v.B)
}

func StructMethod() {
    v := &Ver{3,4}

    fmt.Println(v.Abs())
    
}


/*
可以对包中的 任意 类型定义任意方法，而不仅仅是针对结构体。

但是，不能对来自其他包的类型或基础类型定义方法。
*/
type MyFloat float64

func (f MyFloat) Abs1() float64 {
    if f<0 {
        return float64(-f)
    }
    return float64(f)
}

func StructMethod1() {
    f := MyFloat(-math.Sqrt2)
    fmt.Println(f.Abs1())
}
