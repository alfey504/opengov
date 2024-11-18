package effects

import "github.com/alfey504/opengov/models"

func Inverse(img models.RGBAImage) models.RGBAImage {
	return img.Apply(func(r models.RGBA) models.RGBA {
		return r.Apply(func(u uint8) uint8 {
			return 255 - u
		})
	})
}

func GreyScale(img models.RGBAImage) models.RGBAImage {
	return img.Apply(func(r models.RGBA) models.RGBA {
		greyScaleVal := (0.299 * float64(r.A))
		greyScaleVal += (0.587 * float64(r.G))
		greyScaleVal += (0.114 * float64(r.B))
		greyScale := uint8(greyScaleVal)
		return models.CreateRGBA(greyScale, greyScale, greyScale, r.A)
	})
}

func Sepia(img models.RGBAImage) models.RGBAImage {
	return img.Apply(func(rgba models.RGBA) models.RGBA {

		processColor := func(rgba models.RGBA, x, y, z float64) uint8 {
			r, g, b, _ := rgba.RGBA()
			outColor := float64(r)*x + float64(g)*y + float64(b)*z
			return uint8(max(0, min(outColor, 255)) + 0.5)
		}
		newR := processColor(rgba, 0.393, 0.769, 0.189)
		newG := processColor(rgba, 0.349, 0.686, 0.168)
		newB := processColor(rgba, 0.272, 0.534, 0.131)

		return models.CreateRGBA(newR, newG, newB, rgba.A)
	})
}
