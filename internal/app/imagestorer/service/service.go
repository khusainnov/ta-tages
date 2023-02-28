package service

import (
	"github.com/khusainnov/tag/internal/app/imagestorer/repository"
	tapi "github.com/khusainnov/tag/pkg/tages-api"
)

type Store interface {
	UploadImage(image []byte) (*tapi.UploadImageResponse, error)
	ListImage() (*tapi.ListImagesResponse, error)
}

type Service struct {
	Store
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Store: NewStoreService(repo),
	}
}
