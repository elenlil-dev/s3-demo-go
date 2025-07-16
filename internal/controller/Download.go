package controller

import (
	"context"
	"net/http"

	"go.uber.org/zap"
)

func (c *Controller) Download(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	err := c.service.DownloadFiles(ctx, c.config)
	if err != nil {
		w.Write([]byte("Файлы не загружены, произошла ошибка"))
		c.logg.Error("error controller Download", zap.Error(err))
		return
	}

	w.Write([]byte("Файлы загруженны"))
}
