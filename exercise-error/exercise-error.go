package main

import "fmt"

// 给这个方法添加error返回值，让这个方法不支持负数

// 创建一个新的类型，为其implement error interface
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("can't Sqrt negative number: %v", float64(e))
}

// 返回值的类型是抽象值
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	cur := x

	for i := 0; i < 10; i++ {
		cur = (cur + x/cur) / 2
	}

	return cur, nil
}

func main() {

	if result, err := Sqrt(9); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("result => %v\n", result)
	}
}
