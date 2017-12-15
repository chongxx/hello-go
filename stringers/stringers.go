// Stringer
// 一个普遍存在的接口是fmt包中定义的Stringer
// Stringer是一个可以用字符串描述自己的类型。fmt包（还有其他许多包）使用这个进行输出

// fmt.Print* 在输出结构体的时候调用了结构体的String方法
// 这个和Java中的toString类似
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{
		"Dash",
		22,
	}

	z := Person{
		"Zone",
		213,
	}

	fmt.Println(a, z)
}
