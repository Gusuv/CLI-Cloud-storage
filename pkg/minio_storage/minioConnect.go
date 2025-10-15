package minio_storage

import (
	"cloud_storage/config"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func MinioConnect() (*minio.Client, error) {
	cfg := config.Cfg.MinIo

	client, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKeyID, cfg.SecretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		return nil, fmt.Errorf("minio connect error: %v", err)
	}
	fmt.Println("minio connect success")
	return client, nil
}
