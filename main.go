package main

import (
	"fmt"

	"github.com/alfey504/opengov/adjustments"
	"github.com/alfey504/opengov/opengov"
)

func main() {
	img, err := opengov.MakeColorImage("images/1.jpg")
	if err != nil {
		panic(err)
	}

	x, y := img.Size()
	fmt.Printf("height : %d, width: %d \n", x, y)
	newImg := adjustments.Gamma(img, 2.0)
	newImg.SaveImage("output/test_1.jpg")

}
