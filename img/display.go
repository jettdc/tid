package img

import (
	"fmt"
	"github.com/gookit/color"
	"golang.org/x/term"
	"log"
	"math"
)

type Dimensions struct {
	Width  int
	Height int
}

func PrintPixels(pxs [][]Pixel) {
	fmt.Println("\033[2J")
	for _, pxRow := range pxs {
		fmt.Print("\n")
		for _, px := range pxRow {
			color.Print(fmt.Sprintf("<bg=%d,%d,%d> </>", px.R, px.G, px.B))
		}
	}
	fmt.Print("\n")
}

// Terminal dimensions in characters
func GetTerminalDim() Dimensions {
	if !term.IsTerminal(0) {
		log.Fatal("Must run inside a terminal.")
	}
	width, height, err := term.GetSize(0)
	if err != nil {
		log.Fatal("Something went wrong getting terminal dimensions.")
	}

	println("width:", width, "height:", height)
	return Dimensions{width, height}
}

func getPixelForBlock(pxs [][]Pixel) Pixel {
	return pxs[len(pxs)/2][len(pxs[0])/2]
}

// On monospace font, we treat each character slot as a 2x1 matrix of pixels, meaning that if we
// don't want the image to look stretched out, we must combine y pixels to characters, 2 pixels per character
func TransformImage(pxs [][]Pixel, termSize Dimensions) [][]Pixel {
	fmt.Println("TERM", termSize)

	imgHeight := len(pxs)
	imgWidth := len(pxs[0])

	xScale := int(math.Ceil(float64(imgWidth) / float64(termSize.Width)))

	imgHeight = imgHeight - (imgHeight % xScale)
	imgWidth = imgWidth - (imgWidth % xScale)

	terminalImgDisplayDim := Dimensions{
		imgWidth / xScale,
		imgHeight / xScale,
	}

	newPixels := make([][]Pixel, terminalImgDisplayDim.Height/2+1)
	for i := range newPixels {
		newPixels[i] = make([]Pixel, terminalImgDisplayDim.Width)
	}

	// Every other row
	for row := 0; row < int(math.Floor(float64(terminalImgDisplayDim.Height))); row += 2 {
		lowerYBound := row * xScale
		upperYBound := int(math.Min(float64(row*xScale+(xScale*2)), float64(imgHeight)))

		for col := 0; col < terminalImgDisplayDim.Width; col++ {
			lowerXBound := col * xScale
			upperXBound := int(math.Min(float64(col*xScale+xScale), float64(imgWidth)))

			pxGroup := getGroup(pxs, lowerYBound, upperYBound, lowerXBound, upperXBound)
			newPixels[(row / 2)][col] = getPixelForBlock(pxGroup)
		}
	}

	fmt.Println(len(newPixels[0]), len(newPixels))

	return newPixels
}

func getGroup(pxs [][]Pixel, lowerY int, upperY int, lowerX int, upperX int) [][]Pixel {
	xRange := upperX - lowerX
	yRange := upperY - lowerY

	group := make([][]Pixel, yRange)
	for i := range group {
		group[i] = make([]Pixel, xRange)
	}

	for i := 0; i < yRange; i++ {
		for j := 0; j < xRange; j++ {
			group[i][j] = pxs[lowerY+i][lowerX+j]
		}
	}

	return group
}
