package models

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"sync"

	"github.com/alfey504/opengov/utils"
)

type RGBAImage struct {
	vector      [][]RGBA
	len, height int
}

func MakeColorImage(path string) (RGBAImage, error) {
	img, err := utils.LoadImage(path)
	if err != nil {
		fmt.Printf("Error loading file \n")
		return RGBAImage{}, err
	}

	vec := ImageToVector(*img)
	return RGBAImage{
		vector: vec,
		len:    len(vec),
		height: len(vec[0]),
	}, nil
}

func MakeImageFromVector(vec [][]RGBA) RGBAImage {
	return RGBAImage{
		vector: vec,
		len:    len(vec),
		height: len(vec[0]),
	}
}

func (rgbaImg RGBAImage) SaveImage(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	img := rgbaImg.toImage()
	if err := jpeg.Encode(f, img, nil); err != nil {
		return err
	}

	return nil
}

func (col RGBAImage) GetVector() [][]RGBA {
	return col.vector
}

func (col RGBAImage) At(i, j int) RGBA {
	return col.vector[j][i]
}

func (col RGBAImage) Size() (int, int) {
	return col.height, col.len
}

func (img RGBAImage) Apply(operation func(RGBA) RGBA) RGBAImage {
	x, y := img.Size()

	newImg := make([][]RGBA, x)
	for pos := range newImg {
		newImg[pos] = make([]RGBA, y)
	}

	wg := sync.WaitGroup{}
	for i := 0; i < x; i++ {
		wg.Add(1)
		go func(i int) {
			for j := 0; j < y; j++ {
				p := img.At(i, j)
				newImg[i][j] = operation(p)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	return MakeImageFromVector(newImg)
}

func ImageToVector(img image.Image) [][]RGBA {
	size := img.Bounds().Size()
	vec := make([][]RGBA, size.X)
	for pos := range vec {
		vec[pos] = make([]RGBA, size.Y)
	}
	for i := 0; i < size.X; i++ {
		for j := 0; j < size.Y; j++ {
			col, err := RGBAfromColor(img.At(i, j))
			if err != nil {
				panic(err)
			}

			vec[i][j] = col
		}
	}
	return vec
}

func (img RGBAImage) toImage() image.Image {
	x, y := img.Size()

	rect := image.Rect(0, 0, x, y)
	newImg := image.NewRGBA(rect)
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			p := img.At(i, j).ToColor()
			newImg.Set(i, j, p)
		}
	}
	return newImg
}
