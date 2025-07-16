package run

import (
	"s3-demo/s3-demo-go/internal/config"
	"s3-demo/s3-demo-go/internal/controller"
	"s3-demo/s3-demo-go/internal/infastructure/logger"

	fiber "github.com/gofiber/fiber/v2"
)

type App struct {
	fiber *fiber.App
}

func NewApp(conf *config.Config, crl *controller.Controller) *App {
	r := fiber.New()

	app := &App{
		fiber: r,
	}

	app.Router(crl)
	return app
}

func (a *App) Router(crl *controller.Controller) {
	a.fiber.Get("/create", crl.CreateMany)
	a.fiber.Get("/get", crl.Download)
}

func (a *App) Serve(logg *logger.ZapLogger, conf *config.Config) error {
	logg.Info("Сервер запущен")
	return a.fiber.Listen(conf.HttpAddrServer)
}

func (a *App) Shutdown(logg *logger.ZapLogger) error {
	logg.Info("Сервер получил сигнал и прекратил свою работу")
	return a.fiber.Shutdown()
}
