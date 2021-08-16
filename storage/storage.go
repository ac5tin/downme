package storage

import (
	"errors"
)

type Storage interface {
	Upload(content []byte, id string) error
	Download(id string) ([]byte, error)
}

type storageBase struct {
	Storage
}

func NewStorageBase() *storageBase {
	return &storageBase{}
}

func (sb *storageBase) Upload(content []byte, id string) error {
	return errors.New("upload function not implemented")
}

func (sb *storageBase) Download(id string) ([]byte, error) {
	return []byte(""), errors.New("download function not implemented")
}
