package main

import (
	"os"

	"github.com/alfey504/opengov/blend"
	"github.com/alfey504/opengov/opengov"
)

func main() {
	// fn := func(img opengov.ColorImage) opengov.ColorImage {
	// 	newImg := adjustments.Saturation(img, 20)
	// 	return newImg
	// }
	// TestFiles(fn)
	img1, _ := opengov.MakeColorImage("images/3.jpg")
	img2, _ := opengov.MakeColorImage("images/4.jpg")

	blendedImages := blend.Multiply(img1, img2)
	blendedImages.SaveImage("output/blended.jpg")
}

func getFileNames(folder string) ([]string, error) {
	dirEntry, err := os.ReadDir(folder)
	if err != nil {
		return []string{}, err
	}

	files := make([]string, len(dirEntry))
	for pos, entry := range dirEntry {
		println("\nentry name -> " + entry.Name())
		files[pos] = entry.Name()
	}
	return files, nil
}

func TestFiles(operation func(opengov.ColorImage) opengov.ColorImage) {
	files, err := getFileNames("images/")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if file == ".DS_Store" {
			continue
		}
		fileDir := "images/" + file
		image, err := opengov.MakeColorImage(fileDir)
		if err != nil {
			panic(err)
		}
		newImage := operation(image)
		outputDir := "output/" + file
		newImage.SaveImage(outputDir)
	}
}
