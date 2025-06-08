package main

import (
	"os"

	resize "github.com/readytotouch/readytotouch/internal/logo-resize"
)

func main() {
	if len(os.Args) < 3 {
		println("Usage: go run main.go input.jpg output.png")
		return
	}

	inputPath := os.Args[1]
	outputPath := os.Args[2]

	err := resize.Resize(inputPath, outputPath)
	if err != nil {
		panic(err)
	}
}
