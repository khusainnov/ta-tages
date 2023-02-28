package service

import (
	"github.com/khusainnov/tag/internal/app/imagestorer/repository"
	"github.com/khusainnov/tag/internal/app/imagestorer/service/adapters"
	tapi "github.com/khusainnov/tag/pkg/tages-api"
)

type StoreService struct {
	repo repository.Store
}

func NewStoreService(repo repository.Store) *StoreService {
	return &StoreService{repo: repo}
}

func (s *StoreService) UploadImage(image []byte) (*tapi.UploadImageResponse, error) {
	resp, err := s.repo.SaveImage(image)
	if err != nil {
		return &tapi.UploadImageResponse{}, err
	}

	return &tapi.UploadImageResponse{
		Image: adapters.ImageToPb(resp),
	}, nil
}

func (s *StoreService) ListImage() (*tapi.ListImagesResponse, error) {
	list, err := s.repo.ListImage()
	if err != nil {
		return nil, err
	}

	return adapters.ListImageToPb(list), nil
}
