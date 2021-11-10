package gopherloader

import (
	"bytes"
	"image"
	"image/png"
	"log"
	"os"
)

type GopherLoader struct {}

func (c GopherLoader) Load() (bytes.Buffer, error) {
	f, err := os.Open("resources/gopher.png")
	if err != nil {
		return bytes.Buffer{}, err
	}

	defer f.Close()
	decodedImage, _, err := image.Decode(f)

	var buffer bytes.Buffer
	if err := png.Encode(&buffer, decodedImage); err != nil {
		log.Println("unable to encode image")
	}

	return buffer, err
}