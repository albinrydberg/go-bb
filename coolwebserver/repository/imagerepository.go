package repository

import (
	"bytes"
)

const (
	defaultKey = "default"
)

type imageRepository struct {
	repo map[string]bytes.Buffer
}

func New() Repository {
	return imageRepository{
		repo: make(map[string]bytes.Buffer),
	}
}

func NewWithDefault(imageLoader ImageLoader) Repository {
	repo := imageRepository{
		repo: make(map[string]bytes.Buffer),
	}

	if err := repo.Load(defaultKey, imageLoader); err != nil {
		panic(err)
	}

	return repo
}

func (i imageRepository) Put(key string, buffer bytes.Buffer) {
	i.repo[key] = buffer
}

func (i imageRepository) Get(key string) bytes.Buffer {
	imageBuffer, exists := i.repo[key]
	if !exists {
		return i.repo[defaultKey]
	}

	return imageBuffer
}

func (i imageRepository) Load(key string, imageLoader ImageLoader) error {
	image, err := imageLoader.Load()
	if err != nil {
		return err
	}

	i.Put(key, image)

	return nil
}

var _ Repository = imageRepository{}
