package service

import (
	"bytes"
	"context"
	"fmt"
	"image/jpeg"
	"os"
	"path/filepath"
	"s3-demo/s3-demo-go/internal/config"

	"github.com/disintegration/imaging"
)

func (s *service) UploadFiles(ctx context.Context, conf *config.Config) error {

	dir, err := os.Open(conf.DirFiles)
	if err != nil {
		return fmt.Errorf("error open dir in service:%v", err)
	}
	defer dir.Close()

	files, err := dir.ReadDir(-1)
	if err != nil {
		return fmt.Errorf("error dir.ReadDir in service:%v", err)
	}

	for _, file := range files {

		key, data, err := DecodeImage(file.Name(), conf)
		if err != nil {
			return fmt.Errorf("error DecodeImage in service UploadFiles:%v", err)
		}

		err = s.repo.UploadFiles(ctx, conf, key, data)
		if err != nil {
			return fmt.Errorf("error Create service s.repo.CreateMany: %v", err)
		}

	}

	return nil
}

func DecodeImage(nameImage string, conf *config.Config) (string, []byte, error) {
	path := fmt.Sprintf("%s/%s", conf.DirFiles, nameImage)
	file, err := os.Open(path)
	if err != nil {
		return "", nil, fmt.Errorf("error  DecodeImage in Open:%v", err)
	}
	defer file.Close()

	img, err := imaging.Decode(file)
	if err != nil {
		return "", nil, fmt.Errorf("error  DecodeImage in Decode:%v", err)
	}

	size := imaging.Resize(img, 1024, 0, imaging.Lanczos)

	var buf bytes.Buffer
	opts := jpeg.Options{
		Quality: 75,
	}

	err = jpeg.Encode(&buf, size, &opts)
	if err != nil {
		return "", nil, fmt.Errorf("error  DecodeImage in Encode:%v", err)
	}

	relPath := filepath.Base(path)

	return relPath, buf.Bytes(), nil
}
