// channel 是有类型的管道，可以用channel操作符 <- 对其发送或者接收值
// 默认情况下，在另一端准备好之前，发送和接收都会阻塞。这使得 goroutine 可以在没有明确的锁
// 或者竞争变量的情况下进行同步

package main

import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}

	c <- sum
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)

	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}
