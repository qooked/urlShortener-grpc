package httpServer

import (
	"fmt"
	"log/slog"
	"net/http"
	grpc "shortener-web/grpc/clients/api"
	"shortener-web/http/routers"
)

type GRPCclient *grpc.Client

func StartServer(port int) {
	slog.Info("Starting HTTP server", "Port", fmt.Sprintf("%d", port))
	http.ListenAndServe(fmt.Sprintf(":%d", port), routers.Router())
}
