// Go 的返回值可以被命名，这个时候返回值的命名应该具有一定可读性
// 没有参数的return返回各个变量当前的值，这种用法叫做“裸”返回
// 在短的函数中可以这样用，长度函数这样会影响代码可读性，最好不要用

package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum - 2
	y = sum - x
	return
}

func main() {
	fmt.Println(split(10))
}
