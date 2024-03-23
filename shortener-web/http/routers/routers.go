package routers

import (
	"log/slog"
	"net/http"
	"shortener-web/http/handlers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	r.HandleFunc("/notFound", handlers.LoadNotFoundPage)
	r.HandleFunc("/shorten", handlers.ShortenPage)
	r.HandleFunc("/getOriginalURL", handlers.GetOriginalURL)
	r.HandleFunc("/{shortenedURL}", handlers.LoadShortenedPage)
	r.HandleFunc("/", handlers.LoadMainPage)
	staticDir := http.Dir("./ui/static")
	staticHandler := http.StripPrefix("/static/", http.FileServer(staticDir))
	r.Handle("/static/", staticHandler)
	r.PathPrefix("/static/").Handler(staticHandler)
	return r
}
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Request", "Method", r.Method, "URL", r.URL)
		next.ServeHTTP(w, r)
	})
}
