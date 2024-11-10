package models

import "image/color"

type RGBAf64 struct {
	R float64
	G float64
	B float64
	A float64
}

func MakeRGBAf64(R, G, B, A float64) RGBAf64 {
	return RGBAf64{R, G, B, A}
}

func RGBAf64fromRGBA(rgba color.RGBA) RGBAf64 {
	r := float64(rgba.R) / 255.0
	g := float64(rgba.G) / 255.0
	b := float64(rgba.B) / 255.0
	a := float64(rgba.A) / 255.0

	return MakeRGBAf64(r, g, b, a)
}

func (col RGBAf64) ToRGBA() RGBA {
	r := uint8(col.R * 255)
	g := uint8(col.G * 255)
	b := uint8(col.B * 255)
	a := uint8(col.A * 255)
	return RGBA{r, g, b, a}
}

func (col RGBAf64) Apply(op func(float64) float64) RGBAf64 {
	r := op(col.R)
	g := op(col.G)
	b := op(col.B)
	a := op(col.A)
	return RGBAf64{r, g, b, a}
}
