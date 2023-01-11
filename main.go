package main

import (
	"fmt"
	"image"
	_ "image/jpeg" // decode jpeg images
	_ "image/png"  // decode png images
	"os"
	"strconv"
)

func main() {
	imagePath := "path/to/image"

	// Open image file
	file, err := os.Open(imagePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Decode image
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Open a text file for writing
	textFile, err := os.Create("pixels.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer textFile.Close()

	// Get image bounds
	b := img.Bounds()

	// Iterate over pixels
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			// Format the string to match the desired output
			s := "square 1 |> filled (rgb {" + strconv.Itoa(int(r)) + "} {" + strconv.Itoa(int(g)) + "} {" + strconv.Itoa(int(b)) + "}) |> move (" + strconv.Itoa(x) + ", " + strconv.Itoa(y) + "),\n"
			// Write the string to the text file
			_, err := textFile.WriteString(s)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
	fmt.Println("Done!")
}
