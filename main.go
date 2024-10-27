package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"
)

const ASCII_CHARS = "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'. "

func main() {
	input_path := flag.String("i", "", "Path to the input image")
	output_path := flag.String("o", "output.txt", "Path to the output ASCII file")

	fmt.Println(*input_path)
	if *input_path == "" {
		log.Fatal("Please provide an input image path using -i flag.")
	}

	file, err := os.Open(*input_path)

	if err != nil {
		log.Fatal("Could not open the specified file")
	}

	defer file.Close()

	img, _, err := image.Decode(file)

	if err != nil {
		log.Fatal("Could not open the ")
	}

	grayscaled := rgba_to_gray(img)

	outfile, err := os.Create(*output_path)

	if err != nil {
		log.Fatal("Could not create output file" + err.Error())
	} 

	defer outfile.Close()
	png.Encode(outfile, grayscaled)
}

func rgba_to_gray(img image.Image) *image.Gray {
	var (
		bounds = img.Bounds()
		gray   = image.NewGray(bounds)
	)
	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			var rgba = img.At(x, y)
			gray.Set(x, y, rgba)
		}
	}
	return gray
}
