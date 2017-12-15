// 接口类型是由一组方法定义的集合
// 接口类型的值可以存放实现这些方法的任何值

package main

import (
	"fmt"
	"math"
	"os"
)

type Abser interface {
	Abs() float64
	Up() int // MyFloat 和 Vertex 都没实现Up method 所以运行都会报错
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 隐式接口，类型通过显示接口的方法来实现接口，没有显式声明的必要，so not hava keyword 'implements'
// 隐式接口解耦了实现接口的包和定义接口的包互不依赖

// 包 io 定义了 Reader 和 Writer；其实不一定要这么做？

type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type ReadWrite interface {
	Reader
	Writer
}

func main() {
	//	var a Abser
	//	f := MyFloat(-math.Sqrt2)
	//	v := Vertex{3, 4}
	//
	//	a = f
	//	a = v
	//
	//	fmt.Println(a.Abs())
	var w Writer

	// os.Stdout 实现了 Writer
	w = os.Stdout

	fmt.Fprint(w, "Hello, writer\n")
}
