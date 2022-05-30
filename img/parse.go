package img

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

type Pixel struct {
	R int
	G int
	B int
	A int
}

func init() {
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("jpg", "jpg", jpeg.Decode, jpeg.DecodeConfig)
}

func LoadImage(path string) ([][]Pixel, error) {
	if !imageTypeAccepted(path) {
		return nil, errors.New("Invalid filetype.")
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to open image: %s", err.Error()))
	}
	defer f.Close()

	pixels, err := getPixels(f)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to get pixels from image: %s", err.Error()))
	}

	return pixels, nil
}

func imageTypeAccepted(path string) bool {
	return true
}

// Get the bi-dimensional pixel array
func getPixels(file io.Reader) ([][]Pixel, error) {
	img, _, err := image.Decode(file)

	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var pixels [][]Pixel
	for y := 0; y < height; y++ {
		var row []Pixel
		for x := 0; x < width; x++ {
			row = append(row, rgbaToPixel(img.At(x, y).RGBA()))
		}
		pixels = append(pixels, row)
	}

	return pixels, nil
}

func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{int(r / 257), int(g / 257), int(b / 257), int(a / 257)}
}
