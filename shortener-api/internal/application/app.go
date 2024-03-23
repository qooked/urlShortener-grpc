package application

import (
	"log/slog"
	grpcapp "shortener-api/internal/application/grpc"
)

type App struct {
	GRPCserver *grpcapp.App
}

func New(log *slog.Logger, gRPCport int, DBstring string) *App {
	grpcApp := grpcapp.New(log, gRPCport)
	return &App{
		GRPCserver: grpcApp,
	}
}
