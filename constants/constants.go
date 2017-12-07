// 常量可以是字符、字符串、布尔值、数字类型的值
// 常量不能使用 :=

package main

import "fmt"

const pi = 3.14159

func main() {
	const Language string = "English"
	const Truth = true

	fmt.Println("Hello", pi)
	fmt.Println("Hello", Truth)
	fmt.Println("Hello", Language)

}
