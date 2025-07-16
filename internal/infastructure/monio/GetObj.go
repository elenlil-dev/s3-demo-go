package monio

import (
	"context"
	"fmt"
	"s3-demo/s3-demo-go/internal/config"

	"github.com/minio/minio-go/v7"
)

func (m *MinioClient) DownloadFile(ctx context.Context, conf *config.Config, key string) error {
	path := fmt.Sprintf("%s/%s", conf.DirDownloadPath, key)
	err := m.client.FGetObject(ctx, conf.MinioBucketName, key, path, minio.GetObjectOptions{})
	if err != nil {
		return fmt.Errorf("errors GetFile in Storage:%v", err)
	}
	return nil
}
