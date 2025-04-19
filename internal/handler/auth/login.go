package auth

import (
	"net/http"

	"github.com/Abhishek2010dev/Connecta/internal/dto"
	"github.com/Abhishek2010dev/Connecta/internal/handler"
)

func (a *AuthHandler) LoginPage(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"Title":  "Login",
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

func (a *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	payload := handler.ParseAndDecodeForm[dto.LoginUserPayload](w, r, a.decoder)
	if payload == nil {
		return
	}
}
