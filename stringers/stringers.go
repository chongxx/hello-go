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

// 用于保存ip地址
type IPAddr [4]byte

// 返回格式化的字符串
func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
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

	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for k, v := range addrs {
		fmt.Printf("%v: %v\n", k, v)
	}
}
