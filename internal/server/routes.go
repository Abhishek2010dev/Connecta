package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s *Server) RegisterRoutes() http.Handler {
	router := chi.NewRouter()

	fs := http.FileServer(http.Dir("static"))
	router.Handle("/static/*", http.StripPrefix("/static", fs))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World"))
	})

	return router
}
