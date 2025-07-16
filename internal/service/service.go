package service

import (
	"context"
	"s3-demo/s3-demo-go/internal/config"
)

type RepoMonio interface {
	UploadFiles(ctx context.Context, conf *config.Config, fileName string, data []byte) error
	DownloadFile(ctx context.Context, conf *config.Config, key string) error
	GetListObject(ctx context.Context, conf *config.Config) ([]string, error)
}

type service struct {
	repo RepoMonio
}

func NewService(repo RepoMonio) *service {
	return &service{
		repo: repo,
	}
}
