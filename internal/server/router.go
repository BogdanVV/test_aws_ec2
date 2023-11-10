package server

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (*Server) InitRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("qweasd"))
	})
	r.Get("/health", checkHealth)
	return r
}

func checkHealth(w http.ResponseWriter, r *http.Request) {
	responseBody := map[string]string{
		"data": "api is ok",
	}

	responseJSON, err := json.Marshal(responseBody)
	if err != nil {
		http.Error(w, "smth is terribly wrong", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
