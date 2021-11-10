package repository

import "bytes"

type Repository interface {
	Put(key string, buffer bytes.Buffer)
	Get(key string) bytes.Buffer
}

type ImageLoader interface {
	Load() (bytes.Buffer, error)
}
