package main

import (
	"image"
	"image/color"
	"image/png"
	"image/jpeg"
	"os"
	"fmt"
	"path/filepath"
)

func main() {
	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)	
	}

	defer file.Close()

	imageExtension := filepath.Ext(filePath)
	fmt.Println(imageExtension)

	if imageExtension == ".png" {
		imagePng, err := png.Decode(file)
		if err != nil {
			fmt.Println(err)	
		}

		outImg := invertImage(imagePng)
		
		outFile, _ := os.Create("out.jpeg")
		defer outFile.Close()
		png.Encode(outFile, outImg)

		return
	}

	if imageExtension == ".jpeg" || imageExtension == ".jpg" {
		imageJpeg, err := jpeg.Decode(file)
		if err != nil {
			fmt.Println(err)	
		}

		outImg := invertImage(imageJpeg)

		outFile, _ := os.Create("out.jpeg")
		defer outFile.Close()
		jpeg.Encode(outFile, outImg, nil)

		return
	}
}

func invertImage(img image.Image) (*image.RGBA) {
	imgSize := img.Bounds().Size()
	imgRect := image.Rect(0, 0, imgSize.X, imgSize.Y)
	outImg := image.NewRGBA(imgRect)

	for x := 0; x < imgSize.X; x++ {
		for y := 0; y < imgSize.Y; y++ {
			pixelColor := img.At(x, y)
			originalPixelColor := color.RGBAModel.Convert(pixelColor).(color.RGBA)

			redChannel   := uint8(255 - originalPixelColor.R)
			greenChannel := uint8(255 - originalPixelColor.G)
			blueChannel  := uint8(255 - originalPixelColor.B)
			alphaChannel := uint8(originalPixelColor.A)

			invertedPixelColor := color.RGBA{R: redChannel, G: greenChannel, B: blueChannel, A: alphaChannel}
			outImg.Set(x, y, invertedPixelColor)
		}
	}

	return outImg
}
