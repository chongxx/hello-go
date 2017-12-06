// 变量定义可以包含初始值
// 如果初始值是使用表达式，则类型可以省略，变量会从初始值中获得类型

package main

import "fmt"

var x int = 1
var y int = 2

func main() {
	var c, js, python = true, false, "ok~"
	fmt.Println(x, y, c, js, python)
}
