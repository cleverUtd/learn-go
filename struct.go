package demo

import (
	"fmt"
)


/*
一个结构体（ struct ）就是一个字段的集合。
*/
type Vertex struct {
    X int
    Y int
}

func VertexStruct() {
    fmt.Println(Vertex{1,2})
}

func AccessStruct() {
    v := Vertex{1,2}
    fmt.Println(v.X, v.Y)
}

//结构体字段可以通过结构体指针来访问. 通过指针间接的访问是透明的。
func PointerToStruct() {
    v := Vertex{4,5}
    p := &v
    fmt.Println(&p, p.X)
    p.X = 1e9
    fmt.Println(p, v)
}

/*
结构体文法
结构体文法表示通过结构体字段的值作为列表来新分配一个结构体。

使用 Name: 语法可以仅列出部分字段。（字段名的顺序无关。）

特殊的前缀 & 返回一个指向结构体的指针。
*/
var (
    v1 = Vertex{1,2}  // 类型为 Vertex
    v2 = Vertex{X:1} // Y:0 被省略
    v3 = Vertex{} // X:0 和 Y:0
    p = &Vertex{1,2} // 类型为 *Vertex
)

func StructLiterals() {
    fmt.Println(v1, v2, v3, p)
}
