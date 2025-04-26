package auth

import (
	"log"
	"net/http"

	"github.com/Abhishek2010dev/Go-Htmx-Auth-Example/internal/dto"
	"github.com/Abhishek2010dev/Go-Htmx-Auth-Example/internal/handler"
	"github.com/gorilla/csrf"
)

func (h *AuthHandler) RegisterPage(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"Title":  "Register",
		"Form":   dto.CreateUserPayload{},
		"Errors": map[string]string{},

		csrf.TemplateTag: csrf.TemplateField(r),
	}
	h.renderer.Render(
		w,
		data,
		"layout.html",
		"pages/auth/register.html",
		"components/register-form.html",
	)
}

func (h *AuthHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	payload := handler.ParseAndDecodeForm[dto.CreateUserPayload](w, r, h.decoder)
	if payload == nil {
		return
	}

	if validationError := payload.Validate(h.validate); validationError != nil {
		data := map[string]any{
			"Form":           payload,
			"Errors":         validationError,
			csrf.TemplateTag: csrf.TemplateField(r),
		}
		h.renderer.RenderTemplate(w, "register-form", data, "components/register-form.html")
		return
	}

	emailExists, usernameExists, err := h.userRepository.CheckEmailAndUsername(
		payload.Email,
		payload.Username,
	)
	if err != nil {
		log.Println(err)

		handler.RedirectToErrorPage(w, handler.ErrorResponse{
			Title:   "Server Error",
			Message: "Couldn't verify account details. Please try again.",
		})
		return
	}

	if usernameExists || emailExists {
		errors := map[string]string{}
		if usernameExists {
			errors["username"] = "Username is already taken"
		}

		if emailExists {
			errors["email"] = "Email is already registered"
		}

		data := map[string]any{
			"Form":           payload,
			"Errors":         errors,
			csrf.TemplateTag: csrf.TemplateField(r),
		}
		h.renderer.RenderTemplate(w, "register-form", data, "components/register-form.html")
		return
	}

	payload.Password, err = h.passwordService.HashPassword(payload.Password)
	if err != nil {
		log.Println(err)
		handler.RedirectToErrorPage(w, handler.ErrorResponse{
			Title:   "Server Error",
			Message: "Failed to process password. Try again.",
		})
		return
	}

	userID, err := h.userRepository.Create(payload)
	if err != nil {
		log.Println(err)
		handler.RedirectToErrorPage(w, handler.ErrorResponse{
			Title:   "Server Error",
			Message: "Couldn't create account. Please try again.",
		})
		return
	}

	h.setCookie(w, userID)
	w.Header().Set("HX-Redirect", "/")
}
