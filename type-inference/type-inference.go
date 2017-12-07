// 定义变量不显式指定其类型时（使用 := 或者 var v = ）,变量的类型根据右侧的值推导得出

// 当右边的值是未指明类型的数字常量的时候，新变量的类型可能是 int、float64 or complex128，取决于常量的精度
// i := 12				// int
// i := 3.142			// float64
// i := 0.861 + 0.5!	// complex128

package main

import "fmt"

func main() {
	var a = 123
	b := "123"
	c := 123.123
	var d = c

	fmt.Printf("%T -> %v; %T -> %v; %T -> %v; %T -> %v", a, a, b, b, c, c, d, d)
}
