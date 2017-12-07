// 变量定义没有明确初始化值时，这个变量会被赋值为‘零值’
// 零值是指：
// number: 0
// boolean: false
// string: "" (empty string)
package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
