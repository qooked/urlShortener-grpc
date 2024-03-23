package grpc

import (
	"context"
	"errors"
	"log/slog"
	"time"

	grpclog "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	api "github.com/qooked/url-shortener-proto/generate/go/shortener_api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	api api.ShortenerClient
}

func New(ctx context.Context, log *slog.Logger, address string, timeout time.Duration, retriesCount int) (*Client, error) {
	retryOpts := []grpcretry.CallOption{
		grpcretry.WithCodes(codes.NotFound, codes.Aborted, codes.DeadlineExceeded),
		grpcretry.WithMax(uint(retriesCount)),
		grpcretry.WithPerRetryTimeout(timeout),
	}

	logOpts := []grpclog.Option{
		grpclog.WithLogOnEvents(grpclog.PayloadReceived, grpclog.PayloadSent),
	}

	cc, err := grpc.DialContext(ctx, address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(grpclog.UnaryClientInterceptor(InterceptorLogger(log), logOpts...),
			grpcretry.UnaryClientInterceptor(retryOpts...)))
	if err != nil {
		return nil, err
	}

	return &Client{api: api.NewShortenerClient(cc)}, nil
}

func InterceptorLogger(l *slog.Logger) grpclog.Logger {
	return grpclog.LoggerFunc(func(ctx context.Context, lvl grpclog.Level, msg string, fields ...any) {
		l.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}

func (c *Client) ShortenURL(ctx context.Context, url string) (string, error) {
	response, err := c.api.ShortenURL(ctx, &api.ShortenURLRequest{URL: url})
	if err != nil {
		return "", err
	}

	// Проверяем, является ли ответ сообщением об ошибке
	if errResponse := response.GetError(); errResponse != "" {
		// Если является, возвращаем сообщение об ошибке
		return "", errors.New(errResponse)
	}

	// Если сообщение об ошибке не было получено, возвращаем сокращенный URL
	return response.GetShortenedURL(), nil
}
