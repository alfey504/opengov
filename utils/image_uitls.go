package utils

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func LoadImage(path string) (*image.Image, error) {
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
