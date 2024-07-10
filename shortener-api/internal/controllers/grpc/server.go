package grpc

import (
	"fmt"
	gen "github.com/qooked/url-shortener-proto/generate/go/shortener_api"
	"google.golang.org/grpc"
	"net"
	"shortener-api/pkg/logger"
)

type Server struct {
	gen.UnimplementedShortenerServer
	loggerCreator *logger.Logger
	opts          ServerOptions
}

type ServerOptions struct {
	gen.UnimplementedShortenerServer
	usecase    Usecase
	gRPCserver *grpc.Server
	host       string
	port       int
	baseURL    string
}

func NewServer(loggerCreator *logger.Logger, usecase Usecase, host string, port int) *Server {
	return &Server{
		loggerCreator: loggerCreator,
		opts: ServerOptions{
			usecase:    usecase,
			gRPCserver: grpc.NewServer(),
			host:       host,
			port:       port,
			baseURL:    fmt.Sprintf("%s:%d", host, port),
		},
	}

}

func (s *Server) Run() error {
	listener, err := net.Listen("tcp", s.opts.baseURL)
	if err != nil {
		err = fmt.Errorf("net.Listen(...): %w", err)
		return err
	}

	err = s.opts.gRPCserver.Serve(listener)
	if err != nil {
		err = fmt.Errorf("gRPCserver.Serve(...): %w", err)
		return err
	}

	return nil
}

func (s *Server) ShutDown() {
	s.opts.gRPCserver.Stop()
}
