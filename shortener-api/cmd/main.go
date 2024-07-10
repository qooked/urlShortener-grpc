package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"shortener-api/config"
	cache "shortener-api/internal/adapters/cachedrepository"
	repo "shortener-api/internal/adapters/repository"
	"shortener-api/internal/controllers/grpc"
	uc "shortener-api/internal/usecase"
	"shortener-api/pkg/logger"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigCh := make(chan os.Signal, 1)
	doneCh := make(chan struct{}, 1)

	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	log := logger.NewLogger()
	cfg, err := config.Load()
	if err != nil {
		err = fmt.Errorf("config.Load(): %w", err)
		log.Error(err)
		return
	}

	postgresConn, err := repo.InitPostgres(
		ctx,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Database,
	)
	if err != nil {
		err = fmt.Errorf("repository.InitPostgres(...): %w", err)
		log.Error(err)
		return
	}

	repository := repo.NewRepository(postgresConn)
	redisConn := cache.InitRedis(cfg.Redis.Port, cfg.Redis.Host)
	cahcedRepository := cache.NewCachedRepository(redisConn, repository)
	usecase := uc.NewUsecase(cfg.Url, cahcedRepository)
	srv := grpc.NewServer(log, usecase, cfg.GRPCserver.Host, cfg.GRPCserver.Port)

	go func() {
		srv.Run()
	}()
	select {
	case <-sigCh:
		log.Info("Shutting down...")
	case <-doneCh:
		log.Info("Server shutdown completed")
	}
}
