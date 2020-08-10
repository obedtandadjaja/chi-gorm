package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// API health endpoint
	r.Get("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	fmt.Println("Now serving server at :8000")
	http.ListenAndServe(":8000", r)
}
