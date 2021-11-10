package repository

import (
	"bytes"
)

type imageRepository struct {
	repo map[string]bytes.Buffer
}

func New() Repository {
	return imageRepository{
		repo: make(map[string]bytes.Buffer),
	}
}

func NewPreLoaded(key string, imageLoader ImageLoader) Repository {
	repo := imageRepository{
		repo: make(map[string]bytes.Buffer),
	}

	image, err := imageLoader.Load()
	if err != nil {
		panic(err)
	}

	repo.Put(key, image)

	return repo
}

func (i imageRepository) Put(key string, buffer bytes.Buffer) {
	i.repo[key] = buffer
}

func (i imageRepository) Get(key string) bytes.Buffer {
	return i.repo[key]
}

var _ Repository = imageRepository{}
