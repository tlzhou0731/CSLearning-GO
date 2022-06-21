package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	weight int
	length int
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}
func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.weight, img.length)
}
func (img Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), 255, 255}
}

func main() {
	m := Image{150, 150}
	pic.ShowImage(m)
}
