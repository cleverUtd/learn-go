package demo

import (
	"fmt"
)

/*
map 

map 映射键到值

map 在使用之前必须用 make 来创建；值为 nil 的 map 是空的，并且不能对其赋值。
*/

type Vert struct {
    Lat, Long float64
}

var m map[string] Vert

func MapDemo() {
    m = make(map[string]Vert)

    m["Bell Labs"] = Vert {40.68433, -74.39967}

    fmt.Println(m)
}


/*
修改map
在 map m 中插入或修改一个元素： m[key] = elem

获得元素：elem = m[key]

删除元素：delete(m, key)

通过双赋值检测某个键存在：elem, ok = m[key]。 如果 key 在 m 中， ok 为 true。否则， ok 为 false，并且 elem 是 map 的元素类型的零值。
*/
func MutatingMap() {
    m := make(map[string]int)

    m["answer"] = 42
    fmt.Println("the value:", m["answer"])

    m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

    delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

    v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}