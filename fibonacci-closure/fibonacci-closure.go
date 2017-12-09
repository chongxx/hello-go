package main

import "fmt"

// exercise golang closure, do a fibonacci

func fibonacci() func() int {
	source1 := 0
	source2 := 1

	return func() int {
		fb := source1 + source2
		source1 = source2
		source2 = fb
		return fb
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(f())
	}
}
