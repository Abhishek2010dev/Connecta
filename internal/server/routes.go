package server

import "net/http"

func (s *Server) RegisterRoutes() *http.ServeMux {
	router := http.NewServeMux()
	return router
}
