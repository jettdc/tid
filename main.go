package main

import (
	"github.com/jettdc/tid/img"
	"log"
	"os"
)

func main() {
	path := os.Args[1]

	pixels, err := img.LoadImage(path)
	if err != nil {
		log.Fatal("Could not load image. Error:", err.Error())
	}

	termSize := img.GetTerminalDim()
	transformed := img.TransformImage(pixels, termSize)

	img.PrintPixels(transformed)
}
