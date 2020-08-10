package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	v1 "github.com/obedtandadjaja/chi-gorm/controllers/v1"
)

func InitRouter() *chi.Mux {
	r := chi.NewRouter()

	// middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// // request timeout
	// r.Use(middleware.Timeout(60 * time.Second))

	// API health endpoint
	r.Get("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// api/v1/user
	r.Get("/api/v1/users/{user_id}", v1.GetUser)

	return r
}
