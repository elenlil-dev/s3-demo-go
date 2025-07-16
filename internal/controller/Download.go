package controller

import (
	"net/http"

	"go.uber.org/zap"
)

func (c *Controller) Download(w http.ResponseWriter, r *http.Request) {

	err := c.service.DownloadFiles(c.ctx, c.config)
	if err != nil {
		w.Write([]byte("Файлы не загружены, произошла ошибка"))
		c.logg.Error("error controller Download", zap.Error(err))
		return
	}

	w.Write([]byte("Файлы загруженны"))
}
