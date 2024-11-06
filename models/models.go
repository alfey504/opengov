package models

import "image/color"

type HSL struct {
	h float64
	s float64
	l float64
}

func CreateHSL(h float64, s float64, l float64) HSL {
	return HSL{
		h: h,
		s: s,
		l: l,
	}
}

func (hsl HSL) HSL() (float64, float64, float64) {
	return hsl.h, hsl.s, hsl.l
}

func (hsl HSL) ToRGBA() color.RGBA {
	if hsl.s == 0 {
		v := uint8(hsl.l * 255)
		return color.RGBA{v, v, v, 255}
	}

	hueToRGB := func(p, q, t float64) float64 {
		if t < 0 {
			t += 1
		}

		if t > 1 {
			t -= 1
		}
		if t < 1.0/6.0 {
			return p + (q-p)*6*t
		}
		if t < 1.0/2.0 {
			return q
		}
		if t < 2.0/3.0 {
			return p + (q-p)*(2.0/3.0-t)*6
		}
		return p
	}

	q := 0.0
	if hsl.l < 0.5 {
		q = hsl.l * (1 + hsl.s)
	} else {
		q = hsl.l + hsl.s - hsl.l*hsl.s
	}

	p := 2*hsl.l - q
	hsl.h = hsl.h / 360.0

	r := hueToRGB(p, q, hsl.h+1.0/3.0)
	g := hueToRGB(p, q, hsl.h)
	b := hueToRGB(p, q, hsl.h-1.0/3.0)

	return color.RGBA{
		R: uint8(r * 255),
		G: uint8(g * 255),
		B: uint8(b * 255),
		A: 255,
	}
}
