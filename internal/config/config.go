package config

import "go.uber.org/zap"

type AppMode string

const (
	LocalAppMode AppMode = "local"
)

type Config struct {
	L        *zap.Logger
	AppMode  AppMode `env:"APP_MODE" envDefault:"local"`
	GRPCAddr string  `env:"GRPC_ADDR" envDefault:":9000"`
	HTTPAddr string  `env:"HTTP_ADDR" envDefault:":8082"`
	Path     string  `env:"STORE_PATH" envDefault:"./store/"`
}
