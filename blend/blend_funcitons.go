package blend

import (
	"github.com/alfey504/opengov/models"
)

/* needs work */

func Add(img1, img2 models.RGBAImage) models.RGBAImage {
	fn := func(col1, col2 models.RGBA) models.RGBA {
		newCol := col1.Combine(col2, func(c1, c2 uint8) uint8 {
			return (c1 + c2) / 2
		})
		return newCol
	}
	return Blend(img1, img2, fn)
}

func Multiply(img1, img2 models.RGBAImage) models.RGBAImage {
	fn := func(col1, col2 models.RGBA) models.RGBA {
		return col1.Combine(col2, func(u1, u2 uint8) uint8 {
			return uint8(((uint16(u1) * uint16(u2)) + 175) / 255)
		}).Clamp()
	}
	return Blend(img1, img2, fn)
}

func ColorBurn(img1, img2 models.RGBAImage) models.RGBAImage {

	fn := func(col1, col2 models.RGBA) models.RGBA {
		col1f64, col2f64 := col1.ToRGBAf64(), col2.ToRGBAf64()
		var newR, newG, newB, newA float64
		if col2.R == 0 {
			newR = 0
		} else {
			newR = 1.0 - ((1.0 - col1f64.R) / col2f64.R)
		}

		if col2.G == 0 {
			newG = 0
		} else {
			newG = 1.0 - ((1.0 - col1f64.G) / col2f64.G)
		}

		if col2.B == 0 {
			newB = 0
		} else {
			newB = 1.0 - ((1.0 - col1f64.B) / col2f64.B)
		}

		if col2.A == 0 {
			newA = 0
		} else {
			newA = 1.0 - ((1.0 - col1f64.A) / col2f64.A)
		}

		return models.MakeRGBAf64(newR, newG, newB, newA).ToRGBA()
	}
	return Blend(img1, img2, fn)
}

func Subtract(img1, img2 models.RGBAImage) models.RGBAImage {
	colFn := func(a, b uint8) uint8 {
		return a - b
	}
	fn := func(col1, col2 models.RGBA) models.RGBA {
		return col1.Combine(col2, colFn)
	}
	return Blend(img1, img2, fn)
}

func Divide(img1, img2 models.RGBAImage) models.RGBAImage {
	return Blend(img1, img2, func(r1, r2 models.RGBA) models.RGBA {
		return r2.Combine(r2, func(u1, u2 uint8) uint8 {
			if u2 == 0 {
				return u1
			}
			return (u1 / u2) * 255
		})
	})
}
