package main

import (
	"image"
	"image/jpeg"
	"os"

	"github.com/nfnt/resize"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func transformImage(path string, height, width uint) {
	reader, err := os.Open(path)
	CheckError(err)
	defer reader.Close()
	img, _, err := image.Decode(reader)
	CheckError(err)

	newImg := resize.Resize(width, height, img, resize.NearestNeighbor)
	out, err := os.Create(path + "random_name.jpg")
	CheckError(err)

	err = jpeg.Encode(out, newImg, &jpeg.Options{Quality: 100})
	CheckError(err)
}

func main() {
	transformImage("andrea.jpeg", 1080, 1920)
}
