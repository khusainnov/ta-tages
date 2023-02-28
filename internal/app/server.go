package app

import (
	"context"
	"log"
	"net"

	"github.com/khusainnov/tag/internal/app/imagestorer/endpoint"
	"github.com/khusainnov/tag/internal/app/imagestorer/repository"
	"github.com/khusainnov/tag/internal/app/imagestorer/service"
	"github.com/khusainnov/tag/internal/config"
	"github.com/khusainnov/tag/internal/http"
	tapi "github.com/khusainnov/tag/pkg/tages-api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Run(ctx context.Context, cfg *config.Config) error {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(logUnaryInterceptor(cfg.L)),
		grpc.StreamInterceptor(logStreamInterceptor(cfg.L)),
	)

	if cfg.AppMode == config.LocalAppMode {
		reflection.Register(s)
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	l, err := net.Listen("tcp", cfg.GRPCAddr)
	if err != nil {
		log.Fatalf("failed to listen tcp %s, %v", cfg.GRPCAddr, err)
	}

	httpServer := http.New(cfg)
	httpServer.Start()

	repo := repository.NewRepository(cfg.Path)
	srv := service.NewService(repo)
	e := endpoint.NewEndpoint(srv)

	tapi.RegisterImageServiceServer(s, e)

	cfg.L.Info("starting listening grpc server", zap.Any("PORT", cfg.GRPCAddr))
	if err := s.Serve(l); err != nil {
		cfg.L.Fatal("error service grpc server", zap.Error(err))
	}

	return nil
}

func logUnaryInterceptor(log *zap.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (_ interface{}, err error) {
		log.Info("unary interceptor", zap.Any("method", info.FullMethod))

		return handler(ctx, req)
	}
}

func logStreamInterceptor(log *zap.Logger) grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		log.Info("stream interceptor", zap.Any("method", info.FullMethod))

		return nil
	}
}
