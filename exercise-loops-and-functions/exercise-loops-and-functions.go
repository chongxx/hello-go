// 使用牛顿迭代法开方
// 数学差还是不行啊-_—，算法实现是search出来，解释也没看懂
// 算法答案在这里 https://www.zhihu.com/question/20690553

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {

	cur := x

	for i := 0; i < 10; i++ {
		cur = (cur + x/cur) / 2
	}

	return cur
}

func main() {
	var s float64 = 14823

	fmt.Println("牛顿迭代法开方 => ", Sqrt(s))
	fmt.Println("Golang math.Sqrt开方 => ", math.Sqrt(s))

}
