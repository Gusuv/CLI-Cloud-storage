package repository

import (
	"context"
	"github.com/minio/minio-go/v7"
)

type FileRepo struct {
	mi *minio.Client
}

var Bucket string = "test"

func NewFileRepo(mi *minio.Client) *FileRepo {
	return &FileRepo{mi: mi}
}

func (repo *FileRepo) AddFile(ctx context.Context, path string, user_id int64) error {

	return nil
}
