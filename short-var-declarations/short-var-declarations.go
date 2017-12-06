// 在func中，:= 简洁赋值语句在明确类型的地方使用可以代替var定义
// 函数外的每个语句都必须是关键字开始（package, var, func, etc…), := 结构不能在func外使用

package main

import "fmt"

func main() {
	var i, j = 1, 2
	k := 3
	c, python, java := false, true, "used"
	fmt.Println(i, j, k, c, python, java)
}
