package main

import "fmt"

// Go 的函数定义，看起来没什么不妥，特点是参数类型定义在参数的后面, haha~

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(10, 19))
}
