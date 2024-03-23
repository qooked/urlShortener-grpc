package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	grpc "shortener-web/grpc/clients/api"
	httpServer "shortener-web/http/server"
	"shortener-web/internal/config"
	"time"
)

func main() {
	cfg := config.Get()
	log := setupLogger(cfg.Env)
	timeout := 5 * time.Second
	err := grpc.New(context.Background(), log, fmt.Sprintf("%s:%d", cfg.Addr, cfg.GRPC.Port), timeout, 3)
	if err != nil {
		panic(err)
	}
	slog.Info("Connected to grpc server", slog.String("addr", fmt.Sprintf("%s:%d", cfg.Addr, cfg.GRPC.Port)))
	httpServer.StartServer(cfg.Port)
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case "dev":
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	case "prod":
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case "local":
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	default:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	}
	return log
}
