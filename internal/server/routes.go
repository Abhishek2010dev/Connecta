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
	return stack(router)
}
