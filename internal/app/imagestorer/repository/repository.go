package repository

import (
	"github.com/khusainnov/tag/internal/model"
)

type Store interface {
	SaveImage(image []byte) (*model.Image, error)
	ListImage() ([]*model.Image, error)
}

type Repository struct {
	Store
}

func NewRepository(path string) *Repository {
	return &Repository{
		Store: NewStoreRepository(path),
	}
}
