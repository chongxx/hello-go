package main // 函数也可以作为一个值
import (
	"fmt"
	"math"
)

// 函数同样可以传递，作为函数的参数，作为函数的返回值

// 这个函数接受一个函数作为参数，执行这个参数函数，并给这个参数函数传入(3, 4)的参数
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(3, 4))

	fmt.Println(compute(hypot))

	// math.Pow(x,y float64) float64 -> x**y 即x的y次方
	fmt.Println(compute(math.Pow))
}
