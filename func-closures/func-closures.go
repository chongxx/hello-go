package main

import "fmt"

// closures 闭包是什么？其实就是函数的内部又有一个函数，这个内部函数可以访问上级函数的变量
// 我的理解是Java的Class是类似的 http://www.ruanyifeng.com/blog/2009/08/learning_javascript_closures.html
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	f1 := adder()
	f2 := adder()

	for i := 0; i < 20; i++ {
		fmt.Println(f1(i), f2(2*i))
	}

}
