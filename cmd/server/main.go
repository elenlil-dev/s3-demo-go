package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	"s3-demo/s3-demo-go/internal/config"
	"s3-demo/s3-demo-go/internal/controller"
	"s3-demo/s3-demo-go/internal/infastructure/logger"
	"s3-demo/s3-demo-go/internal/infastructure/monio"
	"s3-demo/s3-demo-go/internal/run"
	"s3-demo/s3-demo-go/internal/service"
	"syscall"

	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()

	logg, err := logger.NewZapLogger()
	if err != nil {
		logg.Error("error in main logger", zap.Error(err))
	}
	cnf, err := config.NewConfig(logg)
	if err != nil {
		logg.Error("error in main config", zap.Error(err))
	}

	repo, err := monio.NewMinioClient(ctx, cnf)
	if err != nil {
		logg.Fatal("error init monio:%v", zap.Error(err))
	}
	svc := service.NewService(repo)
	crl := controller.NewController(svc, cnf, logg, ctx)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	s := run.NewApp(cnf, *crl)

	go func() {
		err := s.Serve(logg)
		if err != nil && err != http.ErrServerClosed {
			logg.Error("error in main Serve() gorutine", zap.Error(err))
		}
	}()
	<-stop

	err = s.Shutdown(ctx, logg)
	if err != nil {
		logg.Error("error in main Shutdown", zap.Error(err))
	}
}
