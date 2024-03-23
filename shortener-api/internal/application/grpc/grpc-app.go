package grpcapp

import (
	"fmt"
	"log/slog"
	"net"
	grpcServer "shortener-api/grpc/server"

	"google.golang.org/grpc"
)

type App struct {
	log        *slog.Logger
	gRPCserver *grpc.Server
	port       int
}

func New(log *slog.Logger, port int) *App {
	gRPC := grpc.NewServer()
	grpcServer.RegisterShortenerServer(gRPC)
	return &App{
		log:        log,
		gRPCserver: gRPC,
		port:       port,
	}

}

func (a *App) Run() error {
	const op = "grpcapp.Run"
	log := a.log.With(
		slog.Int("port", a.port),
	)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("grpc server starting", slog.String("addr", l.Addr().String()))

	if err := a.gRPCserver.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
