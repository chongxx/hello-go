package main

import "fmt"

// 定义函数时如果值有多个，又全都是同一个类型，那么最后一个参数写类型就够了

func swap(x, y, z string) (string, string, string) {
	return z, y, x
}

func main() {
	a, b, c := swap("吗", "好", "你")
	fmt.Println(a, b, c)
}
