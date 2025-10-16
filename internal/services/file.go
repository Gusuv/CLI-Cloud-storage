package services

import (
	"cloud_storage/internal/domain"
	"context"
)

type FileService struct {
	repo domain.FileRepo
}

func NewFileService(repo domain.FileRepo) *FileService {
	return &FileService{repo: repo}
}

func (svc *FileService) AddFile(ctx context.Context, path string, user_id int64) error { return nil }
