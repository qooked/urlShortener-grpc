package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"shortener-tg/bot"
	grpc "shortener-tg/grpc/clients/api"
	"shortener-tg/internal/config"
	"time"
)

func main() {
	cfg := config.Get()
	log := setupLogger(cfg.Env)
	timeout := 5 * time.Second
	client, err := grpc.New(context.Background(), log, "localhost:50051", timeout, 3)
	if err != nil {
		fmt.Println(err)
	}
	bot.Start(cfg.BotToken, client, context.Background(), timeout)
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
