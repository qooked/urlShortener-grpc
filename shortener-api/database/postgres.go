package database

import (
	"database/sql"
	"log/slog"

	_ "github.com/lib/pq"
)

var Postgres *sql.DB

func Init(dbString string) error {
	db, err := sql.Open("postgres", dbString)
	if err != nil {
		return err
	}
	Postgres = db
	slog.Info("connected to postgres")
	return nil
}
