package service

import (
	"context"
	"fmt"
	"os"
	"s3-demo/s3-demo-go/internal/config"
)

func (s *service) DownloadFiles(ctx context.Context, conf *config.Config) error {

	info, err := os.Stat(conf.DirDownloadPath)
	if err != nil {
		return fmt.Errorf("error Stat in service DownloadFiles:%v ", err)
	}
	if !info.IsDir() {
		err := os.MkdirAll(conf.DirDownloadPath, 0777)
		if err != nil {
			return fmt.Errorf("error MkdirAll in service DownloadFiles")
		}
	}

	keys, err := s.repo.GetListObject(ctx, conf)
	if err != nil {
		return fmt.Errorf("errors GetListObject in service GetFiles:%v", err)
	}

	for _, key := range keys {
		err := s.repo.DownloadFile(ctx, conf, key)
		if err != nil {
			return fmt.Errorf("errors GetFile in service GetFiles:%v", err)
		}
	}
	return nil
}
