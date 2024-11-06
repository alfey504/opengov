package blend

import (
	"image/color"

	"github.com/alfey504/opengov/opengov"
	"github.com/alfey504/opengov/utils"
)

func Add(img1, img2 opengov.ColorImage) opengov.ColorImage {
	fn := func(col1, col2 color.RGBA) color.RGBA {
		newR := (col1.R / 2) + (col2.R / 2)
		newG := (col1.G / 2) + (col2.G / 2)
		newB := (col1.B / 2) + (col2.B / 2)
		newA := (col1.A / 2) + (col2.A / 2)

		newCol := color.RGBA{newR, newG, newB, newA}
		return newCol
	}
	return Blend(img1, img2, fn)
}

func Multiply(img1, img2 opengov.ColorImage) opengov.ColorImage {
	fn := func(col1, col2 color.RGBA) color.RGBA {
		return utils.ApplyRGB(col1, col2, func(u1, u2 uint8) uint8 {
			return uint8((uint16(u1)*uint16(u2) + 127) / 255)
		})
	}
	return Blend(img1, img2, fn)
}

func Subtract(img1, img2 opengov.ColorImage) opengov.ColorImage {
	colFn := func(a, b uint8) uint8 {
		return a - b
	}
	fn := func(col1, col2 color.RGBA) color.RGBA {
		return utils.ApplyRGB(col1, col2, colFn)
	}
	return Blend(img1, img2, fn)
}
