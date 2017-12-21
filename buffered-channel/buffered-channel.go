// channel 是可以带缓冲的，make提供第二个参数作为缓冲长度来初始化一个缓冲channel

// 向带缓冲的 channel 发送数据的时候，只有在缓冲区满的时候才会阻塞，而当缓冲区为空的时候，
// 接收操作会阻塞

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)
	ch <- 2
	ch <- 1
	// 下面的操作使缓冲区被填满，发生了deadlock……
	// fatal error: all goroutines are asleep - deadlock!
	// 所以我该怎么理解？？？ 黑人？？？
	v, ok := <-ch
	fmt.Println("haha => ", v, ok)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
