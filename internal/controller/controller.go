package controller

import (
	"context"
	"s3-demo/s3-demo-go/internal/config"
	"s3-demo/s3-demo-go/internal/infastructure/logger"
)

type Service interface {
	UploadFiles(ctx context.Context, conf *config.Config) error
	DownloadFiles(ctx context.Context, conf *config.Config) error
}

type Controller struct {
	service Service
	config  *config.Config
	logg    *logger.ZapLogger
}

func NewController(service Service, config *config.Config, logg *logger.ZapLogger) *Controller {
	return &Controller{
		service: service,
		config:  config,
		logg:    logg,
	}
}
