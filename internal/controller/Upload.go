package controller

import (
	"context"
	"net/http"

	"go.uber.org/zap"
)

func (c *Controller) CreateMany(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	err := c.service.UploadFiles(ctx, c.config)
	if err != nil {
		w.Write([]byte("Файлы не загружены в Монио, произошла ошибка"))
		c.logg.Error("error c.service.Create in controller", zap.Error(err))
		return
	}
	w.Write([]byte("Файлы загружены в Монио"))
}
