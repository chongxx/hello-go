package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	// 创建结构体
	v := Vertex{1, 2}
	fmt.Println(v.X)

	v.X = 1024
	fmt.Println(v.X)

	// 通过指针间接访问结构体字段
	p := &v
	p.X = 1e9
	fmt.Println(v)

	var (
		v1 = Vertex{1, 2}
		v2 = Vertex{Y: 100} // X:0 is implicit
		v3 = Vertex{}       // X:0 and Y:0
		p1 = &Vertex{1, 2}  // has tyep *Vertex
	)
	fmt.Println(v1, v2, v3, p1)
}
