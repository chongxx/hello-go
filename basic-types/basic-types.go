// go 基础数据类型
//  1. bool
//  2. string
//  3. int   int8   int16   int32   int64   uint   uint8   uint32   uint64   uintptr
//  4. byte // uint8 的别名
//  5. rune // int32 的别名， 代表一个unicode码
//  6. float32   float64
//	7. complex64  complex128

package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {

	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
	// 字符串匹配%T对应了变量类型，%v对应了变量值
	// bool(false)
	// uint64(18446744073709551615)
	// complex128((2+3i))
}
