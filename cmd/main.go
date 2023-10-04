package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
	"ueransim-api/internal/server"
	"ueransim-api/internal/service"
	"ueransim-api/internal/transport/rest"
	"ueransim-api/pkg/logger"
)

func init() {
	logger.SetLogger(logger.NewZapLogger())
}

func main() {

	var (
		_, cancel = context.WithCancel(context.Background())
		quit      = make(chan os.Signal, 1)
	)

	var (
		executor              = service.NewExecutorService()
		baseStationService    = service.NewBaseStationService(executor)
		mobileTerminalService = service.NewMobileTerminalService(executor)
	)
	s := service.NewService(baseStationService, mobileTerminalService)
	h := rest.NewHandler(s)
	engine := h.InitRoutes()

	srv := server.New(8080, engine)

	go func() {
		if err := srv.Run(); err != nil {
			logger.Fatal(err)
		}
	}()

	logger.Info("admin started")

	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	sig := <-quit

	ctxSrvStop, cancelSrvStop := context.WithTimeout(context.Background(), 5*time.Second)

	if err := srv.Stop(ctxSrvStop); err != nil {
		logger.Fatal(err)
	}

	cancelSrvStop()
	cancel()

	logger.Infof("admin shutdown signal %v", sig)
}
