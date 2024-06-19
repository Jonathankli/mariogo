package pixel

import "image/color"

type Pixel struct {
	X int
	Y int
	C color.RGBA
}

func AddOffset(pixels []Pixel, x int, y int) []Pixel {
	for i := range pixels {
		pixels[i].X += x
		pixels[i].Y += y
	}

	return pixels
}
