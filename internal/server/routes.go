package server

import (
	"net/http"

	"github.com/Abhishek2010dev/Connecta/internal/handler"
	"github.com/Abhishek2010dev/Connecta/internal/middleware"
	"github.com/Abhishek2010dev/Connecta/internal/renderer"
)

func (s *Server) RegisterRoutes() http.Handler {
	stack := middleware.CreateStack(
		middleware.Logging,
	)
	router := http.NewServeMux()
	renderer := renderer.New("templates")

	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	router.Handle("/", handler.NewHomeHandler(renderer))

	return stack(router)
}
