package main

import (
	"flag"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"math"
	"os"
)

const ASCII_CHARS = "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'. "

func main() {
	input_path := flag.String("i", "", "Path to the input image")
	output_path := flag.String("o", "output.txt", "Path to the output ASCII file")

	flag.Parse()

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
		log.Fatal("Could not decode the specified image")
	}

	grayscale_img := rgba_to_gray(img)

	brightness_map := map_chars_to_brightness(ASCII_CHARS, math.MaxUint8)

	ascii_art := gray_to_ascii(grayscale_img, brightness_map[:])

	output_file, err := os.Create(*output_path)

	if err != nil {
		log.Fatal("Could not create output file" + err.Error())
	}

	defer output_file.Close()

	output_file.WriteString(ascii_art)
}

func map_chars_to_brightness(chars string, max uint8) [256]rune {
	var brightness_table [256]rune
	options_len := len(chars)

	for brightness := 0; brightness < int(max)+1; brightness++ {
		char := int(float64(brightness) / float64(max) * float64(options_len-1))
		brightness_table[brightness] = rune(chars[char])
	}

	return brightness_table
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

func gray_to_ascii(img *image.Gray, mapped_chars []rune) string {
	var ascii_result string
	var bounds = img.Bounds()

	for y := 0; y < bounds.Max.Y; y += 2 {
		for x := 0; x < bounds.Max.X; x++ {
			var pixel = img.GrayAt(x, y)
			ascii_result += string(mapped_chars[pixel.Y])
		}
		ascii_result += "\n"
	}

	return ascii_result
}
