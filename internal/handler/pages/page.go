package pages

import (
	"database/sql"
	"net/http"

	"github.com/Abhishek2010dev/Connecta/internal/middleware"
	"github.com/Abhishek2010dev/Connecta/internal/renderer"
	"github.com/gorilla/mux"
)

type Pages struct {
	renderer      renderer.Renderer
	authMiddlware middleware.Auth
}

func NewPages(renderer renderer.Renderer, db *sql.DB) *Pages {
	return &Pages{
		renderer:      renderer,
		authMiddlware: *middleware.NewAuth(db),
	}
}

func (p *Pages) RegisterRoutes(router *mux.Router) {
	router.Use(p.authMiddlware.RequireAuth)
	router.HandleFunc("/", p.Home).Methods(http.MethodGet)
}
