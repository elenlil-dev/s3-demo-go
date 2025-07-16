package monio

import (
	"bytes"
	"context"
	"fmt"
	"s3-demo/s3-demo-go/internal/config"

	"github.com/minio/minio-go/v7"
)

func (m *MinioClient) UploadFiles(ctx context.Context, conf *config.Config, fileName string, data []byte) error {
	_, err := m.client.PutObject(ctx, conf.MinioBucketName, fileName, bytes.NewReader(data), int64(len(data)), minio.PutObjectOptions{})
	if err != nil {
		return fmt.Errorf("error UploadFiles in infrastructure:%v", err)
	}
	return nil
}
