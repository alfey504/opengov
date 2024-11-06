package blend

import (
	"image/color"

	"github.com/alfey504/opengov/opengov"
)

func Blend(
	img1, img2 opengov.ColorImage,
	blendFunction func(color.RGBA, color.RGBA) color.RGBA,
) opengov.ColorImage {
	x1, y1 := img1.Size()
	x2, y2 := img2.Size()

	x := min(x1, x2)
	y := min(y1, y2)

	newImg := make([][]color.RGBA, x)
	for pos := range newImg {
		newImg[pos] = make([]color.RGBA, y)
	}

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			col1 := img1.At(i, j)
			col2 := img2.At(i, j)

			blendCol := blendFunction(col1, col2)
			newImg[i][j] = blendCol
		}
	}

	return opengov.MakeImageFromVector(newImg)

}
