package main

import (
	"github.com/lordjoaosouza/k3s-concurrency/config"
	"github.com/lordjoaosouza/k3s-concurrency/logger"
	"github.com/lordjoaosouza/k3s-concurrency/server"
	"github.com/lordjoaosouza/k3s-concurrency/server/controller"
	"github.com/lordjoaosouza/k3s-concurrency/service"
	"github.com/rs/zerolog"
	"os"
	"strconv"
	"time"
)

func main() {
	cfg := config.Get()
	logger := logger.New(cfg)

	svc := service.New(cfg, logger)
	ctrl := controller.New(svc, logger)

	svr := server.New(cfg, logger, ctrl)

	logger.Info().Msg(strconv.Itoa(cfg.InternalConfig.ServerPort))
	if err := svr.Start(); err != nil {
		end(logger, err, "failed to start server")
	}
}

func end(logger *zerolog.Logger, err error, message string) {
	logger.Fatal().Err(err).Msgf("%+v: %+v", message, err)
	time.Sleep(time.Millisecond * 50)

	os.Exit(1)
}
