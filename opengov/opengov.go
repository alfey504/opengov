package opengov

import (
	"fmt"
	"image/color"
	"image/jpeg"
	"os"
)

type ColorImage struct {
	vector      [][]color.RGBA
	len, height int
}

// type GreyImage struct {
// 	vector      [][]uint8
// 	len, height int
// }

func MakeColorImage(path string) (ColorImage, error) {
	img, err := loadImage(path)
	if err != nil {
		fmt.Printf("Error loading file \n")
		return ColorImage{}, err
	}

	vec := imageToVector(*img)
	return ColorImage{
		vector: vec,
		len:    len(vec),
		height: len(vec[0]),
	}, nil
}

func MakeImageFromVector(vec [][]color.RGBA) ColorImage {
	return ColorImage{
		vector: vec,
		len:    len(vec),
		height: len(vec[0]),
	}
}

func (col ColorImage) SaveImage(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	img := vectorToImage(col)
	if err := jpeg.Encode(f, img, nil); err != nil {
		return err
	}

	return nil
}

func (col ColorImage) GetVector() [][]color.RGBA {
	return col.vector
}

func (col ColorImage) At(i, j int) color.RGBA {
	return col.vector[j][i]
}

func (col ColorImage) Size() (int, int) {
	return col.height, col.len
}

func (img ColorImage) Apply(operation func(color.RGBA) color.RGBA) ColorImage {
	x, y := img.Size()

	newImg := make([][]color.RGBA, x)
	for pos := range newImg {
		newImg[pos] = make([]color.RGBA, y)
	}

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			p := img.At(i, j)
			newImg[i][j] = operation(p)
		}
	}
	return MakeImageFromVector(newImg)
}
