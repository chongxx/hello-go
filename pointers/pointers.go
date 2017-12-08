// 这就是通常所说的“间接引用”或“非间接引用”?
package main

import "fmt"

func main() {
	i, j := 1024, 2017

	p := &i         // & 符号生成一个指向其对象的指针
	fmt.Println(*p) // 通过pointer读取 i
	*p = 2048       // 通过指针设置 i 的值
	fmt.Println(i)

	p = &j
	*p = *p / 10
	fmt.Println(j)
}
