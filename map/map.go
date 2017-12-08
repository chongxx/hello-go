package main

import "fmt"

// map，Golang的键值对
// map在使用之前必须用make来创建，值为nil的map是空的，并且不能对其赋值

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)

	m["Bell Labs"] = Vertex{
		40.123, -23.921,
	}

	m["a"] = Vertex{
		123, 345,
	}

	n := map[string]Vertex{
		"a": Vertex{
			999, 888,
		},
		"b": Vertex{
			666, 333,
		},
	}

	fmt.Println(m)
	fmt.Println(n)

	// 如果顶级类型只是一个类型名，可以在元素里省略类型
	o := map[string]Vertex{
		"a": {32, 21},
		"b": {12, 11},
	}
	fmt.Println(o)

	// map 的增删查改
	p := make(map[string]int)
	// 增
	p["a"] = 1
	p["b"] = 2
	fmt.Println(p)
	// 查
	fmt.Println(p["a"])
	elem, ok := p["c"] // 如果查询的key在map中，则返回对应的元素和true，如果不存在则返回map元素类型的零值和false
	fmt.Println("elem => ", elem, "ok => ", ok)
	// 删
	delete(p, "a")
	fmt.Println(p)
}
