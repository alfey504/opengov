package models

import (
	"fmt"
	"image/color"
	"math"
)

type RGBA struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

func CreateRGBA(r, g, b, a uint8) RGBA {
	return RGBA{r, g, b, a}
}

func (rgba RGBA) RGBA() (uint8, uint8, uint8, uint8) {
	return rgba.R, rgba.G, rgba.B, rgba.A
}

func RGBAfromColor(col color.Color) (RGBA, error) {
	if col == nil {
		return RGBA{}, fmt.Errorf("color is nil")
	}
	rgbaCol := color.RGBAModel.Convert(col).(color.RGBA)
	return CreateRGBA(rgbaCol.R, rgbaCol.G, rgbaCol.B, rgbaCol.A), nil
}

func (rgba RGBA) Clamp() RGBA {
	newR := min(255, max(0, rgba.R))
	newG := min(255, max(0, rgba.G))
	newB := min(255, max(0, rgba.B))
	newA := min(255, max(0, rgba.A))
	return CreateRGBA(newR, newG, newB, newA)
}

func (rgba RGBA) Apply(op func(uint8) uint8) RGBA {
	newR := op(rgba.R)
	newG := op(rgba.G)
	newB := op(rgba.B)
	newA := op(rgba.A)

	return RGBA{newR, newG, newB, newA}
}

func (rgba RGBA) Combine(rgba2 RGBA, op func(uint8, uint8) uint8) RGBA {
	newR := op(rgba.R, rgba2.R)
	newG := op(rgba.G, rgba2.G)
	newB := op(rgba.B, rgba2.B)
	newA := op(rgba.A, rgba2.A)
	return CreateRGBA(newR, newG, newB, newA)
}

func (rgba RGBA) ToColor() color.Color {
	r, g, b, a := rgba.RGBA()
	return color.RGBA{
		R: r,
		G: g,
		B: b,
		A: a,
	}
}

func (rgba RGBA) ToRGBAf64() RGBAf64 {
	rf64 := float64(rgba.R) / 255.0
	gf64 := float64(rgba.G) / 255.0
	bf64 := float64(rgba.B) / 255.0
	af64 := float64(rgba.A) / 255.0
	return RGBAf64{rf64, gf64, bf64, af64}
}

func (rgba RGBA) ToHSL() HSL {
	r := float64(rgba.R) / 255.0
	g := float64(rgba.G) / 255.0
	b := float64(rgba.B) / 255.0

	max := math.Max(math.Max(r, g), b)
	min := math.Min(math.Min(r, g), b)
	delta := max - min

	lightness := (max + min) / 2

	if delta <= 0.0 {
		return CreateHSL(0, 0, lightness)
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
	return CreateHSL(hue, saturation, lightness)
}
