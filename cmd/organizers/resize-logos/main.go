package main

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"

	"golang.org/x/image/draw"
)

func main() {
	if len(os.Args) < 3 {
		println("Usage: go run main.go input.jpg output.png")
		return
	}

	inputPath := os.Args[1]
	outputPath := os.Args[2]

	// Відкриваємо файл
	inFile, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	// Декодуємо картинку
	srcImg, _, err := image.Decode(inFile)
	if err != nil {
		panic(err)
	}

	bounds := srcImg.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	// Створюємо нове полотно 112x56
	finalImg := image.NewRGBA(image.Rect(0, 0, 112, 56))

	// Беремо кольори кутів
	topLeft := srcImg.At(bounds.Min.X, bounds.Min.Y)
	topRight := srcImg.At(bounds.Max.X-1, bounds.Min.Y)
	bottomLeft := srcImg.At(bounds.Min.X, bounds.Max.Y-1)
	bottomRight := srcImg.At(bounds.Max.X-1, bounds.Max.Y-1)

	// Малюємо фон простими кольорами (наприклад, ділимо полотно на 4 частини)
	halfW := 112 / 2
	halfH := 56 / 2

	draw.Draw(finalImg, image.Rect(0, 0, halfW, halfH), &image.Uniform{topLeft}, image.Point{}, draw.Src)
	draw.Draw(finalImg, image.Rect(halfW, 0, 112, halfH), &image.Uniform{topRight}, image.Point{}, draw.Src)
	draw.Draw(finalImg, image.Rect(0, halfH, halfW, 56), &image.Uniform{bottomLeft}, image.Point{}, draw.Src)
	draw.Draw(finalImg, image.Rect(halfW, halfH, 112, 56), &image.Uniform{bottomRight}, image.Point{}, draw.Src)

	switch {
	case width == height:
		// квадрат 1x1
		dstSmall := image.NewRGBA(image.Rect(0, 0, 56, 56))
		draw.CatmullRom.Scale(dstSmall, dstSmall.Bounds(), srcImg, bounds, draw.Over, nil)

		offsetX := (112 - 56) / 2
		offsetY := (56 - 56) / 2
		draw.Draw(finalImg, image.Rect(offsetX, offsetY, offsetX+56, offsetY+56), dstSmall, image.Point{}, draw.Over)

	case width == 2*height:
		// прямокутник 2x1
		dstFull := image.NewRGBA(image.Rect(0, 0, 112, 56))
		draw.CatmullRom.Scale(dstFull, dstFull.Bounds(), srcImg, bounds, draw.Over, nil)

		draw.Draw(finalImg, dstFull.Bounds(), dstFull, image.Point{}, draw.Over)

	default:
		panic("Image must be square (1:1) or rectangle (2:1)!")
	}

	// Зберігаємо результат
	outFile, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	if strings.HasSuffix(strings.ToLower(outputPath), ".png") {
		err = png.Encode(outFile, finalImg)
	} else {
		err = jpeg.Encode(outFile, finalImg, &jpeg.Options{Quality: 90})
	}
	if err != nil {
		panic(err)
	}

	println("Done!")
}
