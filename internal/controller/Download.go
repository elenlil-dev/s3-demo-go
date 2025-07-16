package controller

import (
	"context"
	"net/http"
	"s3-demo/s3-demo-go/internal/response"
)

func (c *Controller) Download(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	err := c.service.DownloadFiles(ctx, c.config)
	if err != nil {
		response.ResponseJson(w, http.StatusBadRequest, "Файлы не загружены, произошла ошибка")
		return
	}

	response.ResponseJson(w, http.StatusOK, "Файлы загруженны")
}
