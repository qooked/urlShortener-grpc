package grpc

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	gen "github.com/qooked/url-shortener-proto/generate/go/shortener_api"
	"log/slog"
)

type Usecase interface {
	ShortenURL(ctx context.Context, url string) (string, error)
	GetURL(ctx context.Context, url string) (string, error)
}

func (s *Server) ShortenURL(ctx context.Context, r *gen.ShortenURLRequest) (*gen.ShortenURLResponse, error) {
	slog.Info("request", slog.String("url", r.GetUrl()))
	shortenedURL, err := s.opts.usecase.ShortenURL(ctx, r.GetUrl())
	if err != nil {
		err = fmt.Errorf("usecase.ShortenURL(...): %w", err)
		slog.Error("error", err)
		return &gen.ShortenURLResponse{
			Result: &gen.ShortenURLResponse_Error{
				Error: "Не удалось создать короткую ссылку",
			},
		}, nil
	}
	return &gen.ShortenURLResponse{
		Result: &gen.ShortenURLResponse_ShortenedURL{
			ShortenedURL: shortenedURL,
		},
	}, nil
}

func (s *Server) GetURL(ctx context.Context, r *gen.GetURLRequest) (*gen.GetURLResponse, error) {
	OriginalURL, err := s.opts.usecase.GetURL(ctx, r.GetURL())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &gen.GetURLResponse{
				Result: &gen.GetURLResponse_Error{
					Error: "Такой короткой ссылки не существует",
				},
			}, nil
		}
		err = fmt.Errorf("usecase.GetURL(...): %w", err)
		slog.Error("error", err)
		return nil, err
	}
	return &gen.GetURLResponse{
		Result: &gen.GetURLResponse_OriginalURL{
			OriginalURL: OriginalURL,
		},
	}, nil
}
