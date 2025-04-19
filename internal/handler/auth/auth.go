package auth

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/Abhishek2010dev/Connecta/internal/renderer"
	"github.com/Abhishek2010dev/Connecta/internal/repository"
	"github.com/Abhishek2010dev/Connecta/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
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

func (h *AuthHandler) RegisterRoutes(r *mux.Router) {
	r.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, err := r.Cookie("authToken")
			if !errors.Is(err, http.ErrNoCookie) {
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
			h.ServeHTTP(w, r)
		})
	})

	r.HandleFunc("/register", h.RegisterPage).Methods(http.MethodGet)
	r.HandleFunc("/register", h.RegisterHandler).Methods(http.MethodPost)

	r.HandleFunc("/login", h.LoginPage).Methods(http.MethodGet)
	r.HandleFunc("/login", h.LoginHandler).Methods(http.MethodPost)
}
