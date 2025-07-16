package config

import (
	"fmt"
	"s3-demo/s3-demo-go/internal/infastructure/logger"

	"github.com/caarlos0/env/v10"
)

type Config struct {
	HttpAddrServer      string `env:"HTTP_ADDR_SERVER"`
	MinioEndpoint       string `env:"MINIO_ENDPOINT"`
	MinioRootUser       string `env:"MINIO_ROOT_USER"`
	MinioRootPassword   string `env:"MINIO_ROOT_PASSWORD"`
	MinioBucketName     string `env:"MINIO_BUCKET_NAME"`
	MinioUseSSL         bool   `env:"MINIO_USE_SSL"`
	MinioTimeExpiration int    `env:"FILE_TIME_EXPIRATION"`
	DirFiles            string `env:"DIR_FILES"`
	DirDownloadPath     string `env:"DIR_DOWNLOAD_PATH"`
}

func NewConfig(logg *logger.ZapLogger) (*Config, error) {
	cnf := &Config{}
	err := env.Parse(cnf)
	if err != nil {
		return nil, fmt.Errorf("error parse in config:%v", err)
	}

	logg.Info("Конфиги загружены")
	return cnf, nil
}
