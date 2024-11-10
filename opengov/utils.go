package opengov

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func loadImage(path string) (*image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error opening file at path : %s \n", path)
		return nil, err
	}
	defer file.Close()

	stats, err := file.Stat()
	if err != nil {
		fmt.Printf("Error reading file stats \n")
		return nil, err
	}

	fmt.Printf("FIle loaded \n")
	fmt.Printf("file directory : %s\n", path)
	fmt.Printf("file name : %s\n", stats.Name())
	fmt.Printf("file size : %d\n", stats.Size())

	img, format, err := image.Decode(file)
	if err != nil {
		fmt.Printf("Error decoding the image \n")
		return nil, err
	}

	if !(format == "jpeg" || format == "png") {
		err := fmt.Errorf("unsupported format %s", format)
		fmt.Println(err.Error())
		return nil, err
	}

	return &img, nil
}

func imageToVector(img image.Image) [][]color.RGBA {
	size := img.Bounds().Size()
	vec := make([][]color.RGBA, size.X)
	for pos := range vec {
		vec[pos] = make([]color.RGBA, size.Y)
	}
	for i := 0; i < size.X; i++ {
		for j := 0; j < size.Y; j++ {
			col := colorToRGBA(img.At(i, j))
			vec[i][j] = col
		}
	}
	return vec
}

func vectorToImage(img ColorImage) image.Image {
	x, y := img.Size()

	rect := image.Rect(0, 0, x, y)
	newImg := image.NewRGBA(rect)
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			p := img.At(i, j)
			newImg.Set(i, j, p)
		}
	}
	return newImg
}

func colorToRGBA(col color.Color) color.RGBA {
	if col == nil {
		panic(fmt.Errorf("col is nil"))
	}
	return color.RGBAModel.Convert(col).(color.RGBA)
}
