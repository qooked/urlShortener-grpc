package main

import (
	"log/slog"
	"os"
	"shortener-api/database"
	"shortener-api/internal/application"
	"shortener-api/internal/config"
)

func main() {
	config.Get()
	log := setupLogger(config.CFG.Env)
	err := database.Init(config.CFG.DBstring)
	if err != nil {
		log.Error("failed to connect to database", slog.String("error", err.Error()))
		os.Exit(1)
	}
	app := application.New(log, config.CFG.GRPCConfig.Port, config.CFG.DBstring)
	app.GRPCserver.Run()
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
