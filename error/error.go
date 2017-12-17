// Golang programs express error state with error values.
// 函数通常会返回一个error值，调用他的代码应当判断这个错误值是否等于nil来进行错误处理

package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

// 这个函数返回了MyError的struct，在Print这个结构体的时候直接就调用了Error方法
// 为什么不是String方法被调用？难道因为打印error类型的时候就是直接调用Error方法的输出？
// 问题先留着，接下来再看吧
func run() error {
	return MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
