package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"os"

	"github.com/dhowden/tag"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Input file not specified!")
		os.Exit(1)
	}

	if len(args) >= 2 {
		fmt.Println("Expected only one file!")
		os.Exit(1)
	}

	file, err := os.Open(args[0])
	defer file.Close()
	if err != nil {
		fmt.Println("Error when opening file!")
		os.Exit(1)
	}

	m, err := tag.ReadFrom(file)
	if err != nil {
		fmt.Println("Error when reading from file!")
		os.Exit(1)
	}

	if m.Picture() == nil {
		fmt.Println("This file doesn't have a cover image!")
		os.Exit(1)
	}

	imageBytes := m.Picture().Data
	img, _, err := image.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		fmt.Println("Error when decoding cover image!")
		os.Exit(1)
	}

	out, _ := os.Create("./cover.jpg")
	defer out.Close()

	var opts jpeg.Options
	opts.Quality = 95
	err = jpeg.Encode(out, img, &opts)

	if err != nil {
		fmt.Println("Error when exporting cover!")
		os.Exit(1)
	}

	fmt.Println("Title: ", m.Title())
}
