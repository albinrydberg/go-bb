package dogloader

import (
	"bytes"
	"image"
	"image/jpeg"
	"log"
	"os"
)

type DogLoader struct {}

func (d DogLoader) Load() (bytes.Buffer, error) {
	f, err := os.Open("resources/dog.jpg")
	if err != nil {
		return bytes.Buffer{}, err
	}

	defer f.Close()
	decodedImage, _, err := image.Decode(f)

	var buffer bytes.Buffer
	if err := jpeg.Encode(&buffer, decodedImage, nil); err != nil {
		log.Println("unable to encode image")
	}

	return buffer, err
}
