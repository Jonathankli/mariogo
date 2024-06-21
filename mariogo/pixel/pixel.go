package pixel

import (
	"fmt"
	"image/color"
)

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

func GetRelativePixels(ref Pixel, convert []Pixel) []Pixel {

	var relativePixels []Pixel

	for _, pixel := range convert {
		relativePixels = append(relativePixels, Pixel{X: pixel.X - ref.X, Y: pixel.Y - ref.Y, C: pixel.C})
	}

	return relativePixels
}

func GetAbsolutePixels(ref Pixel, convert []Pixel) []Pixel {

	var absolutePixels []Pixel

	for _, pixel := range convert {
		absolutePixels = append(absolutePixels, Pixel{X: pixel.X + ref.X, Y: pixel.Y + ref.Y, C: pixel.C})
	}

	return absolutePixels
}

func PrintPx(pixels []Pixel) {
	for _, pixel := range pixels {
		fmt.Println("Pixel{X:", pixel.X, ", Y:", pixel.Y, ", C: color.RGBA{", pixel.C.R, ",", pixel.C.G, ",", pixel.C.B, ",", pixel.C.A, "}},")
	}
}
