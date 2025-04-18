package auth

import (
	"database/sql"

	"github.com/Abhishek2010dev/Connecta/internal/renderer"
	"github.com/Abhishek2010dev/Connecta/internal/repository"
	"github.com/Abhishek2010dev/Connecta/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
)

type AuthHandler struct {
	renderer        renderer.Renderer
	passwordService service.Password
	sessionService  service.Session
	userRepository  repository.User
	decoder         *schema.Decoder
	validate        *validator.Validate
	tokenName       string
}

func NewAuthHandler(renderer renderer.Renderer, db *sql.DB) *AuthHandler {
	return &AuthHandler{
		renderer:        renderer,
		passwordService: service.NewPasswordService(),
		sessionService:  service.NewSession(db),
		userRepository:  repository.NewUser(db),
		decoder:         schema.NewDecoder(),
		validate:        validator.New(),
		tokenName:       "authToken",
	}
}

func (a *AuthHandler) RegisterRoutes(r chi.Router) {
	r.Get("/register", a.RegisterPage)
	r.Post("/register", a.RegisterHandler)

	r.Get("/login", a.LoginPage)
}
