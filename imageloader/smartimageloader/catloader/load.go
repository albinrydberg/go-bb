package catloader

import (
	"bytes"
	"image"
	"image/jpeg"
	"log"
	"os"
)

type CatLoader struct {}

func (c CatLoader) Load() (bytes.Buffer, error) {
	f, err := os.Open("resources/cat.jpg")
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
