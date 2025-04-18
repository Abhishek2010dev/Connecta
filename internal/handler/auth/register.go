package auth

import (
	"log"
	"net/http"

	"github.com/Abhishek2010dev/Connecta/internal/dto"
	"github.com/Abhishek2010dev/Connecta/internal/handler"
)

func (a *AuthHandler) RegisterPage(w http.ResponseWriter, r *http.Request) {
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

func (a *AuthHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	payload := handler.ParseAndDecodeForm[dto.CreateUserPayload](w, r, a.decoder)
	if payload == nil {
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

	emailExists, usernameExists, err := a.userRepository.CheckEmailAndUsername(payload.Email, payload.Username)
	if err != nil {
		log.Println(err)

		handler.RedirectToErrorPage(w, handler.ErrorResponse{
			Title:   "Server Error",
			Message: "Something went wrong while checking your account details. Please try again later.",
		})
		return
	}

	if usernameExists || emailExists {
		errors := map[string]string{}
		if usernameExists {
			errors["username"] = "This username is already taken"
		}

		if emailExists {
			errors["email"] = "This email is already registered"
		}

		data := map[string]any{
			"Form":   payload,
			"Errors": errors,
		}
		a.renderer.RenderTemplate(w, "register-form", data, "components/register-form.html")
		return
	}

	payload.Password, err = a.passwordService.HashPassword(payload.Password)
	if err != nil {
		log.Println(err)
		handler.RedirectToErrorPage(w, handler.ErrorResponse{
			Title:   "Server Error",
			Message: "Something went wrong while hashing password. Please try again later.",
		})
		return
	}

	userID, err := a.userRepository.Create(payload)
	if err != nil {
		log.Println(err)
		handler.RedirectToErrorPage(w, handler.ErrorResponse{
			Title:   "Server Error",
			Message: "Something went wrong while create your account. Please try again later.",
		})
		return
	}

	token, err := a.sessionService.GenerateToken(userID)
	if err != nil {
		log.Println(err)
		handler.RedirectToErrorPage(w, handler.ErrorResponse{
			Title:   "Server Error",
			Message: "Something went wrong while generating session token. Please try again later.",
		})
		return
	}

	setCookie(w, a.tokenName, token)
	w.Header().Set("HX-Redirect", "/")

}
