package controller

import (
	"fmt"

	fiber "github.com/gofiber/fiber/v2"
)

func (contrl *Controller) CreateMany(c *fiber.Ctx) error {

	ctx := c.Context()

	err := contrl.service.UploadFiles(ctx, contrl.config)
	if err != nil {
		err := c.JSON(fiber.Map{"message": "Файлы не загружены в Монио, произошла ошибка"})
		if err != nil {
			return fmt.Errorf("error json fiber:%v", err)
		}

	}

	err = c.JSON(fiber.Map{"message": "Файлы загружены в Монио"})
	if err != nil {
		return fmt.Errorf("error json fiber:%v", err)
	}
	return nil
}
