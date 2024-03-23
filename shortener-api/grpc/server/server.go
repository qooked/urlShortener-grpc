package api

import (
	"context"
	"encoding/base64"
	"fmt"
	"hash/crc32"
	"net/url"
	"shortener-api/database"
	"shortener-api/internal/config"

	api "github.com/qooked/url-shortener-proto/generate/go/shortener_api"
	"google.golang.org/grpc"
)

type ServerAPI struct {
	api.UnimplementedShortenerServer
}

func RegisterShortenerServer(gRPC *grpc.Server) {
	api.RegisterShortenerServer(gRPC, &ServerAPI{})
}

func (s *ServerAPI) ShortenURL(ctx context.Context, r *api.ShortenURLRequest) (*api.ShortenURLResponse, error) {
	addr := r.GetUrl()
	_, err := url.ParseRequestURI(addr)

	if err != nil {
		return &api.ShortenURLResponse{
			Result: &api.ShortenURLResponse_Error{
				Error: "Некорректно введенная ссылка",
			},
		}, nil
	}
	var count int
	database.Postgres.QueryRow("SELECT COUNT(original_url) FROM url WHERE original_url = $1", addr).Scan(&count)
	fmt.Println(err)
	if count < 1 {
		shortKey := config.Get().Url + generateShortKey(addr)

		_, err = database.Postgres.Exec("INSERT INTO url (original_url, shortened_url) VALUES ($1, $2)", addr, shortKey)
		if err != nil {

			return &api.ShortenURLResponse{
				Result: &api.ShortenURLResponse_Error{
					Error: "Не удалось создать короткую ссылку",
				},
			}, nil
		}

		return &api.ShortenURLResponse{
			Result: &api.ShortenURLResponse_ShortenedURL{
				ShortenedURL: shortKey,
			},
		}, nil
	}

	var existedURL string

	database.Postgres.QueryRow("SELECT shortened_url FROM url WHERE original_url = $1", addr).Scan(&existedURL)

	return &api.ShortenURLResponse{
		Result: &api.ShortenURLResponse_ShortenedURL{
			ShortenedURL: existedURL,
		},
	}, nil
}

func (s *ServerAPI) GetURL(ctx context.Context, r *api.GetURLRequest) (*api.GetURLResponse, error) {
	var count int
	database.Postgres.QueryRow("SELECT COUNT(original_url) FROM url WHERE shortened_url = $1", r.GetURL()).Scan(&count)
	if count < 1 {
		return &api.GetURLResponse{
			Result: &api.GetURLResponse_Error{
				Error: "Такой короткой ссылки не существует",
			},
		}, nil
	}
	var OriginalURL string
	database.Postgres.QueryRow("SELECT original_url FROM url WHERE shortened_url = $1", r.GetURL()).Scan(&OriginalURL)
	return &api.GetURLResponse{
		Result: &api.GetURLResponse_OriginalURL{
			OriginalURL: OriginalURL,
		},
	}, nil
}

func generateShortKey(url string) string {
	crcHash := crc32.ChecksumIEEE([]byte(url))

	base64Hash := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%d", crcHash)))

	return base64Hash
}
