package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// 自定义一个Image类型必要的实现三个方法
// Bounds -> 返回一个image.Rectangle, e.g. 'image.Rect(0,0,w,h)'
// ColorModel -> 返回 color.RGBAModel
// At -> 返回一个颜色 color.RGBA{r, g, b, a}
type Image struct {
	h int
	w int
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.w, img.h)
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) At(x, y int) color.Color {

	var r, g, b uint8

	// 这样就可以生成一个横条纹的图片了，that's interesting~
	if y%2 == 0 {
		r = 0
		g = 0
		b = 0
	} else {
		r = 255
		g = 255
		b = 255
	}

	return color.RGBA{
		r, g, b,
		uint8(255),
	}
}

func main() {
	m := Image{
		100,
		100,
	}
	pic.ShowImage(m)
}
