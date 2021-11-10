package repository

import "bytes"

type Repository interface {
	Put(key string, buffer bytes.Buffer)
	Get(key string) bytes.Buffer

	Load(key string, imageLoader ImageLoader) error
}

type ImageLoader interface {
	Load() (bytes.Buffer, error)
}
