// var 定义了一个变量列表，和函数的参数列表一样类型在变量名后边
// var 可以定义在package or func 级别
package main

import "fmt"

var java, javascript, python, php bool

func main() {
	var i int
	fmt.Println(i, java, javascript, python, php)
}
