// 类型转换
// 表达式 T(variable) 将值 variable 转换为类型T
// 如果返回值已经确定了类型，那么定义变量时可以省略类型不写

package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	z := math.Sqrt(float64(x*x + y*y))
	i := float32(8.6)
	fmt.Printf("%T, %T, %T, %T\n", x, y, z, i)
}
