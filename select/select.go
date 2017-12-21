// select 语句使得一个 goroutine 在多个通讯操作上等待
// select 会阻塞，直到条件分支中的某个可以继续执行，这时就会执行那个条件分支。
// 当多个都准备好的时候，会随机选择一个

// goroutine, channel, select 这几个概念对我来说都比较难理解
// 先不求甚解把，估计多用一下，突然灵光一闪就明白了……

package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacci(c, quit)
}
