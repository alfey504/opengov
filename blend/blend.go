package blend

import (
	"github.com/alfey504/opengov/models"
)

func Blend(
	img1, img2 models.RGBAImage,
	blendFunction func(models.RGBA, models.RGBA) models.RGBA,
) models.RGBAImage {
	x1, y1 := img1.Size()
	x2, y2 := img2.Size()

	x := min(x1, x2)
	y := min(y1, y2)

	newImg := make([][]models.RGBA, x)
	for pos := range newImg {
		newImg[pos] = make([]models.RGBA, y)
	}

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			col1 := img1.At(i, j)
			col2 := img2.At(i, j)

			blendCol := blendFunction(col1, col2)
			newImg[i][j] = blendCol
		}
	}

	return models.MakeImageFromVector(newImg)

}
