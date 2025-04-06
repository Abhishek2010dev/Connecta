package server

import (
	"net/http"

	"github.com/Abhishek2010dev/Connecta/internal/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	stack := middleware.CreateStack(
		middleware.Logging,
	)
	router := http.NewServeMux()
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	return stack(router)
}
