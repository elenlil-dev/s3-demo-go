package minio

import (
	"context"
	"fmt"
	"s3-demo/s3-demo-go/internal/config"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioClient struct {
	client *minio.Client
}

func NewMinioClient(conf *config.Config) (*MinioClient, error) {
	ctx := context.Background()

	client, err := minio.New(conf.MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(conf.MinioRootUser, conf.MinioRootPassword, ""),
		Secure: conf.MinioUseSSL,
	})
	if err != nil {
		return nil, fmt.Errorf("error minio minio.New:%v", err)
	}

	c := &MinioClient{client: client}

	ok, err := c.client.BucketExists(ctx, conf.MinioBucketName)
	if err != nil {
		return nil, fmt.Errorf("error minio client.BucketExists:%v", err)
	}
	if !ok {
		err := c.client.MakeBucket(ctx, conf.MinioBucketName, minio.MakeBucketOptions{})
		if err != nil {
			return nil, fmt.Errorf("error minio client.MakeBucket:%v", err)
		}
	}

	return c, nil
}
