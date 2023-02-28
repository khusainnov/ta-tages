package main

import (
	"context"

	"github.com/caarlos0/env/v6"
	"github.com/khusainnov/tag/internal/app"
	"github.com/khusainnov/tag/internal/config"
	"go.uber.org/zap"
)

func main() {
	log, _ := zap.NewProduction()

	cfg := &config.Config{
		L: log,
	}

	if err := env.Parse(cfg); err != nil {
		log.Fatal("failed to retrieve env variables", zap.Error(err))
	}

	if err := app.Run(context.Background(), cfg); err != nil {
		log.Error("error running grpc server", zap.Error(err))
	}
}
