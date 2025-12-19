package saver

import (
	"image"
	"image/png"
	"os"
)

type PngSaver struct {
}

func NewPngSaver() *PngSaver {
	return &PngSaver{}
}

func (s *PngSaver) Save(img image.Image, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := png.Encode(f, img); err != nil {
		return err
	}
	return nil
}
