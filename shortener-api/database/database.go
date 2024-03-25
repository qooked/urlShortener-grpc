package database

import (
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

var Postgres *sql.DB

var Redis *redis.Client

func InitRedis(redisPort int, redisHost string) error {
	Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisHost, redisPort),
		Password: "",
		DB:       0,
	})

	slog.Info("connected to redis")

	return nil
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

func PostgresMigrate() error {
	for err := Postgres.Ping(); err != nil; {
		err = Postgres.Ping()
		time.Sleep(500 * time.Millisecond)
	}

	_, err := Postgres.Exec("CREATE TABLE IF NOT EXISTS url (id SERIAL PRIMARY KEY, shortened_url TEXT NOT NULL, original_url TEXT NOT NULL); CREATE UNIQUE INDEX IF NOT EXISTS url_alias_key ON url (shortened_url);CREATE INDEX IF NOT EXISTS idx_alias ON url (shortened_url);")
	if err != nil {
		panic("Error creating table: " + err.Error())
	}
	return nil
}
