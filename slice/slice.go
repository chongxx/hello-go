package main

import (
	"fmt"
	"strings"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 7, 8, 9}
	fmt.Println("s == ", s)

	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] == %d\n", i, s[i])
	}

	// slice 可以包含任意类型的值，包括另一个slice
	// 多维数组…
	game := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	game[0][0] = "X"
	game[2][2] = "O"
	game[2][0] = "X"
	printBoard(game)

	// 对slice切片，简单点说就是获取数组里想x~y的元素，等于Js和Java里的Array.slice(begin, end)
	// 一样的包含begin不包含end
	fmt.Println("s ==> ", s)
	fmt.Println("s[1: 4] ==> ", s[1:4])

	// 省略low则从0开始
	fmt.Println("s[:4] ==> ", s[:4])

	// 省略上标则直接结束
	fmt.Println("s[4:] ==> ", s[4:])

	// slice就像是对array的引用，slice不储存any data，只是描述了底层数组的一部分
	// 修改slice的元素将会彻底修改其底层数组的相应元素
	language := [4]string{
		"Java",
		"Js",
		"Python",
		"Elixir",
	}
	fmt.Println(language)

	lan1 := language[:2]
	lan2 := language[1:3]
	fmt.Println(lan1, lan2)

	lan2[0] = "PHP"
	fmt.Println(lan1, lan2)
	fmt.Println(language)

	// slice 由函数make创建。这会分配一个全是零值的数组并返回一个slice指向这个数组
	a := make([]int, 5)
	printSlice("a => ", a)
	b := make([]int, 0, 5)
	printSlice("b => ", b)
	c := b[:2]
	printSlice("c => ", c)
	d := c[2:5]
	printSlice("d => ", d)

	// slice 的零值是 nil
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("z is nil")
	}

	// add element to slice
	f := []int{}
	printSlice("f => ", f)

	f = append(f, 0)
	printSlice("f => ", f)

	f = append(f, 2)
	printSlice("f => ", f)

	f = append(f, 2, 3, 4, 5)
	printSlice("f => ", f)

	// range: for 循环的迭代模式，可以对 slice 和 map 进行循环迭代
	// 当使用 for 循环遍历一个 slice 时，每次迭代 range 将返回两个值
	// 第一个是当前index，第二个是对应的元素的一个拷贝
	type G struct {
		X int
	}

	g := []G{
		G{1},
		G{2},
	}

	for i, v := range g {
		fmt.Printf("2**%d = %v\n", i, v)
		v.X = 9
		fmt.Printf("2**%d = %v\n", i, v)
	}

	fmt.Println("----------- 在for的迭代里修改结构体的值不会影响原slice里的元素")
	for i, v := range g {
		fmt.Printf("2**%d = %v\n", i, v)
	}

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func printBoard(s [][]string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%s \n", strings.Join(s[i], " "))
	}
}
