package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"

	"github.com/disintegration/imaging"
)

func main() {
	// Specify the path to the raw file
	rawFilePath := "path/to/raw/file.raw"

	// Open the raw file
	rawFile, err := imaging.Open(rawFilePath)
	if err != nil {
		fmt.Println("Failed to open the raw file:", err)
		return
	}
	defer rawFile.Close()

	// Decode the raw file into an image
	rawImage, _, err := image.Decode(rawFile)
	if err != nil {
		fmt.Println("Failed to decode the raw file:", err)
		return
	}

	// Create a new JPG file
	jpgFilePath := "path/to/save/jpg/file.jpg"
	jpgFile, err := os.Create(jpgFilePath)
	if err != nil {
		fmt.Println("Failed to create the JPG file:", err)
		return
	}
	defer jpgFile.Close()

	// Encode the raw image as a JPG image
	err = jpeg.Encode(jpgFile, rawImage, nil)
	if err != nil {
		fmt.Println("Failed to encode the raw image as a JPG:", err)
		return
	}

	fmt.Println("Raw file converted to JPG successfully!")
}
