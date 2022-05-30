package main

import (
	"fmt"
	"github.com/jettdc/tid/img"
	"log"
)

func main() {
	pixels, err := img.LoadImage("./testpng.png")
	if err != nil {
		log.Fatal("Could not load image. Error:", err.Error())
	}

	termSize := img.GetTerminalDim()

	transformed := img.TransformImage(pixels, termSize)

	fmt.Println("TSIZE", len(transformed), len(transformed[0]))

	img.PrintPixels(transformed)
}
