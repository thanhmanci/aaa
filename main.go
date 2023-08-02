package main

import (
	"image"
	"image/png"
	"os"

	_ "github.com/oov/psd"
)

func main() {
	file, err := os.Open("input.psd")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	out, err := os.Create("image.png")
	if err != nil {
		panic(err)
	}
	err = png.Encode(out, img)
	if err != nil {
		panic(err)
	}
}
