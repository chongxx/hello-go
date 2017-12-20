package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// 强制只使用一个逻辑处理器
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i < 100; i++ {
			for j := 0; j < 100000; j++ {
				f := i + j*i*j
				if f == -2 {
					fmt.Print("dsad")
				}
			}
			fmt.Println("A:", i)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 1; i < 100; i++ {
			for j := 0; j < 100000; j++ {
				f := i + j*i*j
				if f == -2 {
					fmt.Print("dsad")
				}
			}
			fmt.Println("B:", i)
		}
	}()
	wg.Wait()
}
