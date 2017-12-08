// defer 表达式会延迟函数的执行直到包裹他的函数执行完才执行
// defer 的函数不会马上执行但是函数的参数是马上生成的
package main

import "fmt"

func main() {
	defer1()
	defer2()
}

func defer1() {
	i := 1
	defer fmt.Println("world", i)

	i = 2
	fmt.Println("hello", i)

	// hello 2
	// world 1
}

// defer 函数的调用被压入一个栈中，当包裹这些defer的方法return时，defer栈的函数会弹出来执行
func defer2() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
