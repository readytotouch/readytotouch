package logo_resize

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/image/draw"
	"golang.org/x/image/webp"
)

func Resize112x56(inputPath string, outputPath string) error {
	inFile, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	var srcImg image.Image
	switch strings.ToLower(filepath.Ext(inputPath)) {
	case ".webp":
		srcImg, err = webp.Decode(inFile)
		if err != nil {
			panic(err)
		}
	default:
		srcImg, _, err = image.Decode(inFile)
		if err != nil {
			panic(err)
		}
	}

	var (
		bounds   = srcImg.Bounds()
		width    = bounds.Dx()
		height   = bounds.Dy()
		finalImg = image.NewRGBA(image.Rect(0, 0, 112, 56))
	)

	switch {
	case width == height:
		var (
			topLeft     = srcImg.At(bounds.Min.X, bounds.Min.Y)
			topRight    = srcImg.At(bounds.Max.X-1, bounds.Min.Y)
			bottomLeft  = srcImg.At(bounds.Min.X, bounds.Max.Y-1)
			bottomRight = srcImg.At(bounds.Max.X-1, bounds.Max.Y-1)
			halfW       = 112 / 2
			halfH       = 56 / 2
		)

		draw.Draw(finalImg, image.Rect(0, 0, halfW, halfH), &image.Uniform{topLeft}, image.Point{}, draw.Src)
		draw.Draw(finalImg, image.Rect(halfW, 0, 112, halfH), &image.Uniform{topRight}, image.Point{}, draw.Src)
		draw.Draw(finalImg, image.Rect(0, halfH, halfW, 56), &image.Uniform{bottomLeft}, image.Point{}, draw.Src)
		draw.Draw(finalImg, image.Rect(halfW, halfH, 112, 56), &image.Uniform{bottomRight}, image.Point{}, draw.Src)

		// квадрат 1x1
		dstSmall := image.NewRGBA(image.Rect(0, 0, 56, 56))
		draw.CatmullRom.Scale(dstSmall, dstSmall.Bounds(), srcImg, bounds, draw.Over, nil)

		offsetX := (112 - 56) / 2
		offsetY := (56 - 56) / 2
		draw.Draw(finalImg, image.Rect(offsetX, offsetY, offsetX+56, offsetY+56), dstSmall, image.Point{}, draw.Over)

	case width == 2*height:
		dstFull := image.NewRGBA(image.Rect(0, 0, 112, 56))
		draw.CatmullRom.Scale(dstFull, dstFull.Bounds(), srcImg, bounds, draw.Over, nil)

		draw.Draw(finalImg, dstFull.Bounds(), dstFull, image.Point{}, draw.Over)

	default:
		return errors.New("unsupported image aspect ratio, only square and 2:1 aspect ratios are supported")
	}

	outFile, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	if strings.HasSuffix(strings.ToLower(outputPath), ".png") {
		err = png.Encode(outFile, finalImg)
		if err != nil {
			return err
		}

		return nil
	}

	err = jpeg.Encode(outFile, finalImg, &jpeg.Options{Quality: 90})
	if err != nil {
		return err
	}

	return nil
}

func Resize72x72(inputPath string, outputPath string) error {
	inFile, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	var srcImg image.Image
	switch strings.ToLower(filepath.Ext(inputPath)) {
	case ".webp":
		srcImg, err = webp.Decode(inFile)
		if err != nil {
			panic(err)
		}
	default:
		srcImg, _, err = image.Decode(inFile)
		if err != nil {
			panic(err)
		}
	}

	var (
		bounds   = srcImg.Bounds()
		width    = bounds.Dx()
		height   = bounds.Dy()
		finalImg = image.NewRGBA(image.Rect(0, 0, 72, 72))
	)

	switch {
	case width == height:
		dstFull := image.NewRGBA(image.Rect(0, 0, 72, 72))
		draw.CatmullRom.Scale(dstFull, dstFull.Bounds(), srcImg, bounds, draw.Over, nil)

		draw.Draw(finalImg, dstFull.Bounds(), dstFull, image.Point{}, draw.Over)

	default:
		return errors.New("unsupported image aspect ratio, only square and 2:1 aspect ratios are supported")
	}

	outFile, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	if strings.HasSuffix(strings.ToLower(outputPath), ".png") {
		err = png.Encode(outFile, finalImg)
		if err != nil {
			return err
		}

		return nil
	}

	err = jpeg.Encode(outFile, finalImg, &jpeg.Options{Quality: 90})
	if err != nil {
		return err
	}

	return nil
}
