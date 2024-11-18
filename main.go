package main

import (
	"os"
	"sync"

	"github.com/alfey504/opengov/models"
)

func main() {
	// img1, _ := models.LoadRGBAImage("images/3.jpg")
	// img2, _ := models.LoadRGBAImage("images/6.png")

	// blendedImages := blend.Multiply(img1, img2)
	// blendedImages.SaveImage("output/blended.jpg")

	TestFiles(func(r models.RGBAImage) models.RGBAImage {
		kernel, _ := models.CreateKernel([][]float64{
			{1, 0, -1},
			{1, 0, -1},
			{1, 0, -1},
		})
		return r.Convolve(kernel)
	})
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

func TestFiles(operation func(models.RGBAImage) models.RGBAImage) {
	files, err := getFileNames("images/")
	if err != nil {
		panic(err)
	}
	wg := sync.WaitGroup{}
	for _, file := range files {
		if file == ".DS_Store" {
			continue
		}
		fileDir := "images/" + file
		image, err := models.LoadRGBAImage(fileDir)
		if err != nil {
			panic(err)
		}

		// newImage := operation(image)
		// outputDir := "output/" + file
		// newImage.SaveImage(outputDir)

		wg.Add(1)
		go func(image models.RGBAImage, file string) {
			newImage := operation(image)
			outputDir := "output/" + file
			newImage.SaveImage(outputDir)
			wg.Done()
		}(image, file)
	}
	wg.Wait()
}
