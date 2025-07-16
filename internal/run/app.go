package run

import (
	"context"
	"net/http"
	"s3-demo/s3-demo-go/internal/config"
	"s3-demo/s3-demo-go/internal/controller"
	"s3-demo/s3-demo-go/internal/infastructure/logger"
	"time"

	"github.com/go-chi/chi"
)

type App struct {
	http   *http.Server
	router *chi.Mux
}

func NewApp(conf *config.Config, crl controller.Controller) *App {
	r := chi.NewRouter()

	app := &App{
		http: &http.Server{
			Addr:    conf.HttpAddrServer,
			Handler: r,
		},
		router: r,
	}

	app.Router(crl)
	return app
}

func (a *App) Router(crl controller.Controller) {
	a.router.Get("/create", crl.CreateMany)
	a.router.Get("/get", crl.Download)
}

func (a *App) Serve(logg *logger.ZapLogger) error {
	logg.Info("Сервер запущен")
	return a.http.ListenAndServe()
}

func (a *App) Shutdown(logg *logger.ZapLogger) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	logg.Info("Сервер получил сигнал и прекратил свою работу")
	return a.http.Shutdown(ctx)
}
