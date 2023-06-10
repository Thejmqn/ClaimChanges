package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"path/filepath"
)

func main() {
	folderPath := "src/github.com/TheJmqn/firstapp/maps"
	files, err := os.ReadDir(folderPath)
	if err != nil {
		log.Fatal(err)
	}

	water := loadImage("src/github.com/TheJmqn/firstapp/Water.png")
	images := make([]image.Image, 0)

	for _, file := range files {
		filePath := filepath.Join(folderPath, file.Name())
		img := loadImage(filePath)
		images = append(images, img)
	}

	pixels := [3252][3252]int{}

	for i := 1; i < len(images); i++ {
		for x := 0; x < images[i].Bounds().Dx(); x++ {
			for y := 0; y < images[i].Bounds().Dy(); y++ {
				prevR, prevG, prevB := colors(x, y, images[i-1])
				curR, curG, curB := colors(x, y, images[i])
				if prevR != curR || prevG != curG || prevB != curB {
					pixels[x][y]++
				}
			}
		}
	}

	writeImage := image.NewRGBA(image.Rect(0, 0, images[0].Bounds().Dx(), images[0].Bounds().Dy()))
	for x := 0; x < writeImage.Bounds().Dx(); x++ {
		for y := 0; y < writeImage.Bounds().Dy(); y++ {
			switch pixels[x][y] {
			case 0:
				writeImage.Set(x, y, color.RGBA{255, 255, 255, 0xff})
			case 1:
				writeImage.Set(x, y, color.RGBA{255, 0, 0, 0xff})
			case 2:
				writeImage.Set(x, y, color.RGBA{255, 128, 0, 0xff})
			case 3:
				writeImage.Set(x, y, color.RGBA{255, 255, 0, 0xff})
			case 4:
				writeImage.Set(x, y, color.RGBA{0, 255, 0, 0xff})
			case 5:
				writeImage.Set(x, y, color.RGBA{0, 0, 255, 0xff})
			default:
				writeImage.Set(x, y, color.RGBA{255, 0, 255, 0xff})
			}
			if isWater(x, y, water, images[0]) {
				writeImage.Set(x, y, color.RGBA{0, 0, 0, 0})
			}
		}
	}

	f, _ := os.Create("image.png")
	png.Encode(f, writeImage)
}

func loadImage(file string) image.Image {
	imageFile, err := os.Open(file)
	if err != nil {
		log.Println("Error opening file:", file, err)
	}
	defer imageFile.Close()

	img, err := png.Decode(imageFile)
	if err != nil {
		log.Println("Error decoding image:", file, err)
	}

	return img
}

func colors(x int, y int, image image.Image) (int, int, int) {
	rgbColor := image.At(x, y)
	r, g, b, _ := rgbColor.RGBA()
	normalizedR := int(r >> 8)
	normalizedG := int(g >> 8)
	normalizedB := int(b >> 8)
	return normalizedR, normalizedG, normalizedB
}

func isWater(x int, y int, water image.Image, anyMap image.Image) bool {
	r, b, g := 55, 79, 106
	colorR, colorB, colorG := colors(x, y, water)
	return (r == colorR && b == colorB && g == colorG)
}
