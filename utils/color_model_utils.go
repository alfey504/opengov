package utils

import (
	"image/color"
	"math"

	"github.com/alfey504/opengov/models"
)

func RGBtoHSL(col color.RGBA) models.HSL {
	r := float64(col.R) / 255.0
	g := float64(col.G) / 255.0
	b := float64(col.B) / 255.0

	max := math.Max(math.Max(r, g), b)
	min := math.Min(math.Min(r, g), b)
	delta := max - min

	lightness := (max + min) / 2

	if delta <= 0.0 {
		return models.CreateHSL(0, 0, lightness)
	}

	saturation := 0.0
	if lightness < 0.5 {
		saturation = delta / (max + min)
	} else {
		saturation = delta / (2.0 - max - min)
	}

	hue := 0.0
	switch max {
	case r:
		hue = (g - b) / delta
	case g:
		hue = 2.0 + ((b - r) / delta)
	case b:
		hue = 4.0 + ((r - g) / delta)
	}

	hue *= 60
	if hue < 0 {
		hue += 360
	}
	return models.CreateHSL(hue, saturation, lightness)
}

func ApplyRGB(col1, col2 color.RGBA, op func(uint8, uint8) uint8) color.RGBA {
	newR := max(min(255, op(col1.R, col2.R)), 0)
	newG := max(min(255, op(col1.G, col2.G)), 0)
	newB := max(min(255, op(col1.B, col2.B)), 0)
	newA := max(min(255, op(col1.A, col2.A)), 0)
	return color.RGBA{newR, newG, newB, newA}
}
