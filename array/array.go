package main

import "fmt"

func main() {
	// array 是定长的
	var a [2]string
	a[0] = "Hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
}
