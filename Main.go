package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("Starting")

	folderPath := "src/github.com/TheJmqn/firstapp/maps"
	files, err := os.ReadDir(folderPath)
	if err != nil {
		log.Fatal(err)
	}

	images := make([]image.Image, 0)

	for _, file := range files {
		filePath := filepath.Join(folderPath, file.Name())
		img := loadImage(filePath)
		images = append(images, img)
	}

	targetR, targetG, targetB := 255, 6, 6
	counter := 0

	for _, image := range images {
		for x := 0; x < image.Bounds().Dx(); x++ {
			for y := 0; y < image.Bounds().Dy(); y++ {
				r, g, b := colors(x, y, images[1])
				if r == targetR && g == targetG && b == targetB {
					counter++
				}
			}
		}
	}
	fmt.Println(counter)
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
