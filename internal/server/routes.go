package server

import (
	"net/http"
	"os"

	"github.com/Abhishek2010dev/Connecta/internal/handler"
	"github.com/Abhishek2010dev/Connecta/internal/handler/auth"
	"github.com/Abhishek2010dev/Connecta/internal/renderer"
	"github.com/gorilla/csrf"
	middleware "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var CSRFTokenKey = "X-CSRF-Token"

func (s *Server) RegisterRoutes() http.Handler {
	router := mux.NewRouter()
	renderer := renderer.New("templates")

	router.Use(middleware.ProxyHeaders)
	router.Use(middleware.RecoveryHandler())
	router.Use(func(h http.Handler) http.Handler {
		return middleware.LoggingHandler(os.Stdout, h)
	})

	fs := http.FileServer(http.Dir("static"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	protected := router.NewRoute().Subrouter()
	protected.Use(csrf.Protect(
		[]byte(s.cfg.Server.CsrfSecure),
		csrf.Secure(false),
		csrf.RequestHeader("X-CSRF-Token"),
	))

	protected.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(CSRFTokenKey, csrf.Token(r))
		renderer.Render(w, nil, "layout.html", "error/404.html")
	})

	protected.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		errResp := handler.ErrorResponse{
			Title:   query.Get("title"),
			Message: query.Get("message"),
		}
		w.Header().Set(CSRFTokenKey, csrf.Token(r))
		renderer.Render(w, errResp, "layout.html", "error/other.html")
	}).Methods(http.MethodGet)

	authHandler := auth.NewAuthHandler(renderer, s.db)
	authHandler.RegisterRoutes(protected.PathPrefix("/auth").Subrouter())

	return router
}
