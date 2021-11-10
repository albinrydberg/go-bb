package imageloader

import (
	"bytes"
	"image"
	"image/jpeg"
	"log"
	"os"
)

func LoadGoat() (*bytes.Buffer, error) {
	f, err := os.Open("resources/goat.jpeg")
	if err != nil {
		return nil, err
	}

	defer f.Close()
	decodedImage, _, err := image.Decode(f)

	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, decodedImage, nil); err != nil {
		log.Println("unable to encode image")
	}

	return buffer, err
}

func LoadCat() (*bytes.Buffer, error) {
	f, err := os.Open("resources/cat.jpg")
	if err != nil {
		return nil, err
	}

	defer f.Close()
	decodedImage, _, err := image.Decode(f)

	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, decodedImage, nil); err != nil {
		log.Println("unable to encode image")
	}

	return buffer, err
}

func LoadDog() (*bytes.Buffer, error) {
	f, err := os.Open("resources/dog.jpg")
	if err != nil {
		return nil, err
	}

	defer f.Close()
	decodedImage, _, err := image.Decode(f)

	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, decodedImage, nil); err != nil {
		log.Println("unable to encode image")
	}

	return buffer, err
}