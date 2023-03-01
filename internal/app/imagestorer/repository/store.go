package repository

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"syscall"
	"time"

	"github.com/google/uuid"
	"github.com/khusainnov/tag/internal/model"
)

type StoreRepository struct {
	Path string
}

func NewStoreRepository(path string) *StoreRepository {
	return &StoreRepository{Path: path}
}

func (s *StoreRepository) SaveImage(image []byte) (*model.Image, error) {
	name := uuid.New().String()

	path := filepath.Join(s.Path, fmt.Sprintf("%s.pdf", name))
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		return nil, fmt.Errorf("cannot create file, %w", err)
	}

	if _, err = file.Write(image); err != nil {
		return nil, fmt.Errorf("cannot write data into file, %w", err)
	}

	info, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("cannot get file info, %w", err)
	}

	return getImageData(info)
}

func (s *StoreRepository) ListImage() ([]*model.Image, error) {
	files, err := os.ReadDir(s.Path)
	if err != nil {
		return nil, fmt.Errorf("cannot read store folder, %w", err)
	}
	resp := make([]*model.Image, len(files))

	for i, file := range files {
		info, err := file.Info()
		if err != nil {
			return nil, fmt.Errorf("cannot get file info, %w", err)
		}
		resp[i], err = getImageData(info)
	}

	return resp, nil
}

func (s *StoreRepository) DownloadImage(id string) ([]byte, error) {
	file, err := os.Open(filepath.Join(s.Path, fmt.Sprintf("%s.pdf", id)))
	if err != nil {
		return nil, fmt.Errorf("cannot get image, %w", err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("cannot read image data, %w", err)
	}

	return data, nil
}

func getImageData(info os.FileInfo) (*model.Image, error) {
	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		return nil, fmt.Errorf("unable to get creation date of file")
	}

	createdAt := time.Unix(stat.Ctimespec.Unix())

	return &model.Image{
		Name:      info.Name(),
		CreatedAt: createdAt,
		EditedAt:  info.ModTime(),
	}, nil
}
