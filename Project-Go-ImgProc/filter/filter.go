package filter

import (
	"fmt"
	"image/jpeg"
	"os"

	"github.com/disintegration/imaging"
)

type Filter interface {
	Process(srcPath, dstPath string) error
}

type Grayscale struct{}

func (g Grayscale) Process(srcPath, dstPath string) error {
	src, err := imaging.Open(srcPath)
	if err != nil {
		fmt.Printf("Failed to open image: %v", err)
		return err
	}

	img := imaging.Grayscale(src)

	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	opts := jpeg.Options{Quality: 90}
	return jpeg.Encode(dstFile, img, &opts)
}

type Blur struct{}

func (b Blur) Process(srcPath, dstPath string) error {
	src, err := imaging.Open(srcPath)
	if err != nil {
		fmt.Printf("Failed to open image: %v", err)
		return err
	}

	img := imaging.Blur(src, 5.0)

	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	opts := jpeg.Options{Quality: 90}
	return jpeg.Encode(dstFile, img, &opts)
}
