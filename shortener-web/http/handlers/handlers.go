package handlers

import (
	"encoding/json"
	"html/template"
	"net/http"

	grpc "shortener-web/grpc/clients/api"
)

func GetOriginalURL(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ParseForm() err: " + err.Error()))
		return
	}

	originalURL, err := grpc.GRPCclient.GetURL(r.Context(), r.FormValue("url"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	type urlResponse struct {
		OriginalURL string `json:"OriginalURL"`
	}

	response := urlResponse{
		OriginalURL: originalURL,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func ShortenPage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ParseForm() err: " + err.Error()))
		return
	}

	shortenedURL, err := grpc.GRPCclient.ShortenURL(r.Context(), r.FormValue("url"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	type urlResponse struct {
		ShortenedURL string `json:"shortened_url"`
	}

	response := urlResponse{
		ShortenedURL: shortenedURL,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func LoadMainPage(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/main-page.html")

	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ts.Execute(w, "")
}

func LoadShortenedPage(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/shortenedURL-page.html")

	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ts.Execute(w, "")
}

func LoadNotFoundPage(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/notFound-page.html")

	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ts.Execute(w, "")
}
