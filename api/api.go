package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/nyumbapoa/backend/internal/config"
)

func New(enableCORS bool, cfg *config.Config) (*chi.Mux, error) {

	r := chi.NewRouter()
	return r, nil
}

func corsConfig() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:     []string{"*"},
		AllowedMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:     []string{"Accept", "Authorization", "Accept-Encoding", "Content-Type", "Content-Length", "X-CSRF-Token"},
		ExposedHeaders:     []string{"Link"},
		AllowCredentials:   true,
		MaxAge:             86400,
		OptionsPassthrough: false,
	})
}
