package controller

import (
	"context"
	"fmt"

	fiber "github.com/gofiber/fiber/v2"
)

func (contrl *Controller) Download(c *fiber.Ctx) error {
	ctx := context.Background()
	c.Accepts("application/json")

	err := contrl.service.DownloadFiles(ctx, contrl.config)
	if err != nil {
		err := c.JSON(fiber.Map{"message": "Файлы не загружены, произошла ошибка"})
		if err != nil {
			return fmt.Errorf("error json fiber:%v", err)
		}
	}

	err = c.JSON(fiber.Map{"message": "Файлы успешно загружены"})
	if err != nil {
		return fmt.Errorf("error json fiber:%v", err)
	}
	return nil
}
