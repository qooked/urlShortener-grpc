package database

import (
	"database/sql"
	"fmt"
	"log/slog"

	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

var Postgres *sql.DB

var Redis *redis.Client

func InitRedis(redisPort int, redisHost string) {
	Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisHost, redisPort),
		Password: "",
		DB:       0,
	})

	slog.Info("connected to redis")
}

func InitPostgres(dbString string) error {
	db, err := sql.Open("postgres", dbString)
	if err != nil {
		return err
	}

	Postgres = db
	slog.Info("connected to postgres")

	return nil
}
