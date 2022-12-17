package main

import (
	"image/color"

	fbuf "github.com/gonutz/framebuffer"
)

func main() {
	fb, err := fbuf.Open("/dev/fb1")
	if err != nil {
		panic(err)
	}
	var c color.Color
	for i := 0; i < 255; i++ {
		fb.Set(i, i, c)
	}
}
