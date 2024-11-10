package adjustments

import (
	"math"

	"github.com/alfey504/opengov/models"
)

func Brightness(img models.RGBAImage, val float64) models.RGBAImage {
	brightnessFactor := (1 + val)

	brightnessFunction := func(val uint8, factor float64) uint8 {
		b := max(0, min(float64(val)*factor, 255))
		return uint8(b)
	}

	fn := func(p models.RGBA) models.RGBA {
		return p.Apply(func(u uint8) uint8 {
			return brightnessFunction(u, brightnessFactor)
		})
	}

	return img.Apply(fn)
}

func Contrast(img models.RGBAImage, val float64) models.RGBAImage {
	contrastFactor := (259 * (val + 255)) / (255 * (259 - val))

	contrastFunction := func(val uint8, factor float64) uint8 {
		c := max(0, min(factor*(float64(val)-128)+128, 255))
		return uint8(c)
	}

	fn := func(p models.RGBA) models.RGBA {
		return p.Apply(func(u uint8) uint8 {
			return contrastFunction(u, contrastFactor)
		})
	}
	return img.Apply(fn)
}

func Gamma(img models.RGBAImage, factor float64) models.RGBAImage {
	gammaCorrection := 1 / factor

	gammaFunction := func(col uint8, factor float64) uint8 {
		g := max(0, min(255*math.Pow((float64(col)/255), factor), 255))
		return uint8(g)
	}

	fn := func(p models.RGBA) models.RGBA {
		return p.Apply(func(u uint8) uint8 {
			return gammaFunction(u, gammaCorrection)
		})
	}
	return img.Apply(fn)
}

func Hue(img models.RGBAImage, factor int) models.RGBAImage {
	fn := func(col models.RGBA) models.RGBA {
		hsl := col.ToHSL()
		h, s, l := hsl.HSL()
		h = float64((int(h) + factor) % 360)
		return models.CreateHSL(h, s, l).ToRGBA()
	}
	return img.Apply(fn)
}

func Saturation(img models.RGBAImage, factor float64) models.RGBAImage {
	fn := func(col models.RGBA) models.RGBA {
		hsl := col.ToHSL()
		h, s, l := hsl.HSL()
		s = max(min(s*(1+factor), 1.0), 0.0)
		newCol := models.CreateHSL(h, s, l).ToRGBA()
		newCol.A = col.A
		return newCol
	}

	return img.Apply(fn)
}
