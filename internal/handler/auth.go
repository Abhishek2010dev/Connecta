package handler

import (
	"database/sql"
	"net/http"

	"github.com/Abhishek2010dev/Connecta/internal/renderer"
	"github.com/Abhishek2010dev/Connecta/internal/repository"
	"github.com/Abhishek2010dev/Connecta/internal/service"
	"github.com/go-chi/chi/v5"
)

type Auth struct {
	renderer        renderer.Renderer
	passwordService service.Password
	sessionService  service.Session
	userRepository  repository.User
}

func NewAuth(renderer renderer.Renderer, db *sql.DB) *Auth {
	return &Auth{
		renderer:        renderer,
		passwordService: service.NewPasswordService(),
		sessionService:  service.NewSession(db),
		userRepository:  repository.NewUser(db),
	}
}

func (a *Auth) RegisterPage(w http.ResponseWriter, r *http.Request) {
	a.renderer.Render(w, map[string]any{
		"Title": "Register",
	}, "pages/auth/layout.html", "pages/auth/register.html")
}

func (a *Auth) RegisterRoutes(r chi.Router) {
	r.Get("/register", a.RegisterPage)
}
