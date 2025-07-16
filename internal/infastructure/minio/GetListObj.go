package minio

import (
	"context"
	"fmt"
	"s3-demo/s3-demo-go/internal/config"

	"github.com/minio/minio-go/v7"
)

func (m *MinioClient) GetListObject(ctx context.Context, conf *config.Config) ([]string, error) {

	var key []string
	ojs := m.client.ListObjects(ctx, conf.MinioBucketName, minio.ListObjectsOptions{})

	for object := range ojs {
		if object.Err != nil {
			return nil, fmt.Errorf("error GetListObject: %v", object.Err)
		}
		key = append(key, object.Key)
	}
	return key, nil
}
