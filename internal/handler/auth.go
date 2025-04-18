package handler

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Abhishek2010dev/Connecta/internal/dto"
	"github.com/Abhishek2010dev/Connecta/internal/renderer"
	"github.com/Abhishek2010dev/Connecta/internal/repository"
	"github.com/Abhishek2010dev/Connecta/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
)

type Auth struct {
	renderer        renderer.Renderer
	passwordService service.Password
	sessionService  service.Session
	userRepository  repository.User
	decoder         *schema.Decoder
	validate        *validator.Validate
}

func NewAuth(renderer renderer.Renderer, db *sql.DB) *Auth {
	return &Auth{
		renderer:        renderer,
		passwordService: service.NewPasswordService(),
		sessionService:  service.NewSession(db),
		userRepository:  repository.NewUser(db),
		decoder:         schema.NewDecoder(),
		validate:        validator.New(),
	}
}

func (a *Auth) RegisterPage(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"Title":  "Register",
		"Form":   dto.CreateUserPayload{},
		"Errors": map[string]string{},
	}
	a.renderer.Render(
		w,
		data,
		"pages/auth/layout.html",
		"pages/auth/register.html",
		"components/register-form.html",
	)
}

func (a *Auth) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Printf("Failed to process request: %s", err)
		error := ErrorResposne{
			Title:   "Something went wrong",
			Message: "We couldn't process your request. Please try again.",
		}
		redirectToErrorPage(w, error)
		return
	}

	var payload dto.CreateUserPayload
	if err := a.decoder.Decode(&payload, r.PostForm); err != nil {
		log.Printf("Invalid submission: %s", err)
		error := ErrorResposne{
			Title:   "Invalid submission",
			Message: "There was an issue with the information you entered. Please review and try again.",
		}
		redirectToErrorPage(w, error)
		return
	}

	if validationError := payload.Validate(a.validate); validationError != nil {
		data := map[string]any{
			"Form":   payload,
			"Errors": validationError,
		}
		a.renderer.RenderTemplate(w, "register-form", data, "components/register-form.html")
		return
	}

	exits, err := a.userRepository.ExistsByEmailAndUsername(payload.Email, payload.Username)
	if err != nil {
		log.Println(err)
		error := ErrorResposne{
			Title:   "Server Error",
			Message: "Something went wrong while checking your account details. Please try again later.",
		}
		redirectToErrorPage(w, error)
	}

	if exits {
		data := map[string]any{
			"Form": payload,
			"Errors": map[string]string{
				"email":    "This email is already registered.",
				"username": "This username is already taken.",
			},
		}
		a.renderer.RenderTemplate(w, "register-form", data, "components/register-form.html")
		return
	}

}

func (a *Auth) RegisterRoutes(r chi.Router) {
	r.Get("/register", a.RegisterPage)
	r.Post("/register", a.RegisterHandler)
}
