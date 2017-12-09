package main

import "fmt"

// Go does not hava classes, However, you can define methods on types.
// A method is an function with a special receiver argument.

// The receiver appears in its own argument list between the func keyword and the method name.

// In this example, the Abs method has a receiver of type Vertex named v

type Vertex struct {
	X, Y float64
}

// 给结构体添加一个method，语法如下，正常的func在func关键字和方法名中间价格(type Type)
// 调用就是直接通过创建这个结构体.mathod()就Ok
func (v Vertex) Mult(i float64) float64 {
	return v.X * v.Y * i
}

func main() {
	v := Vertex{8, 3}
	fmt.Println(v.Mult(10))
}
