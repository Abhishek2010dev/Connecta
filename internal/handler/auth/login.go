package auth

import (
	"net/http"

	"github.com/Abhishek2010dev/Connecta/internal/dto"
)

func (a *AuthHandler) LoginPage(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"Title":  "Register",
		"Form":   dto.CreateUserPayload{},
		"Errors": map[string]string{},
	}
	a.renderer.Render(
		w,
		data,
		"pages/auth/layout.html",
		"pages/auth/login.html",
		"components/login-form.html",
	)
}
