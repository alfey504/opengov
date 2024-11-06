package adjustments

import (
	"image/color"
	"math"

	"github.com/alfey504/opengov/models"
	"github.com/alfey504/opengov/opengov"
	"github.com/alfey504/opengov/utils"
)

func Brightness(img opengov.ColorImage, val float64) opengov.ColorImage {
	brightnessFactor := (1 + val)

	brightnessFunction := func(val uint8, factor float64) uint8 {
		b := max(0, min(float64(val)*factor, 255))
		return uint8(b)
	}

	fn := func(p color.RGBA) color.RGBA {
		return color.RGBA{
			R: brightnessFunction(p.R, brightnessFactor),
			G: brightnessFunction(p.G, brightnessFactor),
			B: brightnessFunction(p.B, brightnessFactor),
			A: brightnessFunction(p.A, brightnessFactor),
		}
	}

	return img.Apply(fn)
}

func Contrast(img opengov.ColorImage, val float64) opengov.ColorImage {
	contrastFactor := (259 * (val + 255)) / (255 * (259 - val))

	contrastFunction := func(val uint8, factor float64) uint8 {
		c := max(0, min(factor*(float64(val)-128)+128, 255))
		return uint8(c)
	}

	fn := func(p color.RGBA) color.RGBA {
		return color.RGBA{
			R: contrastFunction(p.R, contrastFactor),
			G: contrastFunction(p.G, contrastFactor),
			B: contrastFunction(p.B, contrastFactor),
			A: contrastFunction(p.A, contrastFactor),
		}
	}
	return img.Apply(fn)
}

func Gamma(img opengov.ColorImage, factor float64) opengov.ColorImage {
	gammaCorrection := 1 / factor

	gammaFunction := func(col uint8, factor float64) uint8 {
		g := max(0, min(255*math.Pow((float64(col)/255), factor), 255))
		return uint8(g)
	}

	fn := func(p color.RGBA) color.RGBA {
		return color.RGBA{
			R: gammaFunction(p.R, gammaCorrection),
			G: gammaFunction(p.G, gammaCorrection),
			B: gammaFunction(p.B, gammaCorrection),
			A: gammaFunction(p.A, gammaCorrection),
		}
	}
	return img.Apply(fn)
}

func Hue(img opengov.ColorImage, factor int) opengov.ColorImage {
	fn := func(col color.RGBA) color.RGBA {
		hsl := utils.RGBtoHSL(col)
		h, s, l := hsl.HSL()
		h = float64((int(h) + factor) % 360)
		return models.CreateHSL(h, s, l).ToRGBA()
	}
	return img.Apply(fn)
}

func Saturation(img opengov.ColorImage, factor float64) opengov.ColorImage {
	fn := func(col color.RGBA) color.RGBA {
		hsl := utils.RGBtoHSL(col)
		h, s, l := hsl.HSL()
		s = max(min(s*(1+factor), 1.0), 0.0)
		newCol := models.CreateHSL(h, s, l).ToRGBA()
		newCol.A = col.A
		return newCol
	}
	return img.Apply(fn)
}
