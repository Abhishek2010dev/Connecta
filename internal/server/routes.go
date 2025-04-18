package server

import (
	"net/http"
	"os"

	"github.com/Abhishek2010dev/Connecta/internal/handler"
	"github.com/Abhishek2010dev/Connecta/internal/handler/auth"
	"github.com/Abhishek2010dev/Connecta/internal/renderer"
	middleware "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func (s *Server) RegisterRoutes() http.Handler {
	router := mux.NewRouter()
	renderer := renderer.New("templates")

	router.Use(middleware.ProxyHeaders)
	router.Use(middleware.RecoveryHandler())
	router.Use(func(h http.Handler) http.Handler {
		return middleware.LoggingHandler(os.Stdout, h)
	})

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		renderer.Render(w, nil, "layout.html", "error/404.html")
	})

	router.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		error := handler.ErrorResponse{
			Title:   query.Get("title"),
			Message: query.Get("message"),
		}
		renderer.Render(w, error, "layout.html", "error/other.html")
	}).Methods(http.MethodGet)

	fs := http.FileServer(http.Dir("static"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	authHandler := auth.NewAuthHandler(renderer, s.db)
	authHandler.RegisterRoutes(router.PathPrefix("/auth").Subrouter())

	return router
}
