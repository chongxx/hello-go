// Golang只有 for 这一种循环结构
// 和常见的语言类似for包含三个组成部分
// 初始化语句：在第一次循环前被执行（optional）
// 循环条件表达式：每轮迭代开始前被求值
// 后置语句：每轮迭代后背执行(optional)

// 和其他语言最大的区别就是表达式不需要括号()，循环体必须用{}包起来

// init and post statement are optional

package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)

	// 初始化语句和后置语句都是可选的，如果两个都省略其实就是其他语言的 while
	sum = 1
	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)

	// 如果把循环条件也省略掉，就是死循环了
	//	for {
	//	}
}
