package server

import (
	"net/http"

	"github.com/Abhishek2010dev/Connecta/internal/handler"
	"github.com/Abhishek2010dev/Connecta/internal/handler/auth"
	"github.com/Abhishek2010dev/Connecta/internal/renderer"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	router := chi.NewRouter()
	renderer := renderer.New("templates")

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		renderer.Render(w, nil, "layout.html", "error/404.html")
	})

	router.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		error := handler.ErrorResponse{
			Title:   query.Get("title"),
			Message: query.Get("message"),
		}
		renderer.Render(w, error, "layout.html", "error/other.html")
	})

	fs := http.FileServer(http.Dir("static"))
	router.Handle("/static/*", http.StripPrefix("/static", fs))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World"))
	})

	authHandler := auth.NewAuthHandler(renderer, s.db)
	router.Route("/auth", authHandler.RegisterRoutes)

	return router
}
