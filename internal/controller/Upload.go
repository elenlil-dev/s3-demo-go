package controller

import (
	"net/http"

	"go.uber.org/zap"
)

func (c *Controller) CreateMany(w http.ResponseWriter, r *http.Request) {

	err := c.service.UploadFiles(c.ctx, c.config)
	if err != nil {
		w.Write([]byte("Файлы не загружены в Монио, произошла ошибка"))
		c.logg.Error("error c.service.Create in controller", zap.Error(err))
		return
	}
	w.Write([]byte("Файлы загружены в Монио"))
}
