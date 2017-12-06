// Go 的包里需要导出的func名字首字母需要大写，小写开头的func不会被导出
// 我感觉这就是 public & Private
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
}
