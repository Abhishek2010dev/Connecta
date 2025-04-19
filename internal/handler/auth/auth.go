package auth

import (
	"database/sql"
	"errors"
	"net/http"

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
	tokenKey        string
}

func NewAuthHandler(renderer renderer.Renderer, db *sql.DB) *AuthHandler {
	return &AuthHandler{
		renderer:        renderer,
		passwordService: service.NewPasswordService(),
		sessionService:  service.NewSession(db),
		userRepository:  repository.NewUser(db),
		decoder:         schema.NewDecoder(),
		validate:        validator.New(),
		tokenKey:        "authToken",
	}
}

func (a *AuthHandler) RegisterRoutes(r chi.Router) {
	r.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, err := r.Cookie("authToken")
			for !errors.Is(err, http.ErrNoCookie) {
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
			h.ServeHTTP(w, r)
		})
	})

	r.Get("/register", a.RegisterPage)
	r.Post("/register", a.RegisterHandler)

	r.Get("/login", a.LoginPage)
	r.Post("/login", a.LoginHandler)
}
