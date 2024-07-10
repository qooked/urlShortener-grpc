package repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func InitPostgres(ctx context.Context, host string, port int, user string, password string, database string) (*sqlx.DB, error) {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, database)
	db, err := sqlx.ConnectContext(ctx, "postgres", connString)
	if err != nil {
		err = fmt.Errorf("sqlx.Open(...): %w", err)
		return nil, err
	}

	return db, nil
}

func PostgresMigrate() error {
	//for err := Postgres.Ping(); err != nil; {
	//	err = Postgres.Ping()
	//	time.Sleep(500 * time.Millisecond)
	//}
	//
	//_, err := Postgres.Exec("CREATE TABLE IF NOT EXISTS url (id SERIAL PRIMARY KEY, shortened_url TEXT NOT NULL, original_url TEXT NOT NULL); CREATE UNIQUE INDEX IF NOT EXISTS url_alias_key ON url (shortened_url);CREATE INDEX IF NOT EXISTS idx_alias ON url (shortened_url);")
	//if err != nil {
	//	panic("Error creating table: " + err.Error())
	//}
	return nil
}

func (r *Repository) GetURL(ctx context.Context, URL string) (string, error) {
	var originalURL string
	err := r.db.GetContext(ctx, &originalURL, "SELECT original_url FROM url WHERE shortened_url = $1", URL)
	if err != nil {
		err = fmt.Errorf("r.db.Get(...): %w", err)
		return "", err
	}

	return originalURL, nil
}

func (r *Repository) ShortenURL(ctx context.Context, URL string, shortkey string) error {
	_, err := r.db.ExecContext(ctx, `INSERT INTO url (original_url, shortened_url) VALUES ($1, $2)`, URL, shortkey)
	if err != nil {
		err = fmt.Errorf("r.db.Exec(...): %w", err)
		return err
	}
	return nil
}
