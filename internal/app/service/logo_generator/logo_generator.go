package logo_generator

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

type Pixel struct {
	X     int
	Y     int
	Color color.RGBA
}

func Generate(imageName string) {
	imageMatrix := selectImage(imageName)

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	for _, pixel := range imageMatrix {
		for dx := 0; dx < PixelSize; dx++ {
			for dy := 0; dy < PixelSize; dy++ {
				img.Set(pixel.Y*PixelSize+dy, pixel.X*PixelSize+dx, pixel.Color)
			}
		}
	}

	f, _ := os.Create("../assets/images/amazing_logo.png")
	png.Encode(f, img)
}

func selectImage(imageName string) []Pixel {
	switch imageName {
	case "frog":
		return FrogImagePixels
	}
	return nil
}
