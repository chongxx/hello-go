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
	fmt.Println("Mult2 => ", Mult2(v, 10))
	f := MyFloat(-3)
	fmt.Println("Abs => ", f.Abs())
}

// 可以对包中的任意类型定义任意方法，不仅是针对结构体，但是不能对基础类型和其他包的类型定义方法
// You can declare a method on non-struct types, too.
// In this example we see a numeric type MyFloat with a Abs method.
// You can only declare a method with a receiver whose type is defined in the same package as the method.
// You cannot declare a method with a receiver whose type is defined in another package
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// 上里面的例子如果变成常规写法,结果没有任何区别
func Mult2(v Vertex, i float64) float64 {
	return v.X * v.Y * i
}
