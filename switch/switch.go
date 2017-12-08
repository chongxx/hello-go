// Golang switch statement like other language
// 不同的点
// 1. case 语句后面不用加break; Go已经自动提供
// 2. cases 不需要一定是constant

package main

import (
	"fmt"
	"runtime"
)

func main() {
	switch1()
	switch2()
}

func switch1() {
	fmt.Print("Go run on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Mac Os X.")

	case "linux":
		fmt.Println("Linux.")

	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}

// switch 的条件语句不存在时就默认为true，这样可以很干净的写出 if-then-else 链
// case 的参数可以是一个func
func f() int {
	return 1
}

func switch2() {
	switch 1 {
	case 0:
		fmt.Println(0)
	case f():
		fmt.Println("case statement param is function")
	default:
		fmt.Println("not match")
	}
}
