package main

import (
	"fmt"
	"image"
)

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 50, 50))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(2, 2).RGBA())
}
