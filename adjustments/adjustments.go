package adjustments

import (
	"image/color"
	"math"

	"github.com/alfey504/opengov/opengov"
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
	x, y := img.Size()

	newImg := make([][]color.RGBA, x)
	for pos := range newImg {
		newImg[pos] = make([]color.RGBA, y)
	}

}
