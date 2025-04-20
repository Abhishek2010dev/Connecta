package auth

import (
	"log"
	"net/http"
	"strings"

	"github.com/Abhishek2010dev/Connecta/internal/dto"
	"github.com/Abhishek2010dev/Connecta/internal/handler"
	"github.com/gorilla/csrf"
)

func (h *AuthHandler) LoginPage(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"Title":          "Login",
		"Form":           dto.CreateUserPayload{},
		"Errors":         map[string]string{},
		csrf.TemplateTag: csrf.TemplateField(r),
	}
	h.renderer.Render(
		w,
		data,
		"pages/auth/layout.html",
		"pages/auth/login.html",
		"components/login-form.html",
	)
}

func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	payload := handler.ParseAndDecodeForm[dto.LoginUserPayload](w, r, h.decoder)
	if payload == nil {
		return
	}

	if validationError := payload.Validate(h.validate); validationError != nil {
		data := map[string]any{
			"Form":           payload,
			"Errors":         validationError,
			csrf.TemplateTag: csrf.TemplateField(r),
		}
		h.renderer.RenderTemplate(w, "login-form", data, "components/login-form.html")
		return
	}

	userData, err := h.userRepository.FindByEmail(strings.ToLower(payload.Email))
	if err != nil {
		log.Printf("Failed to fetch user by email: %s", err)
		handler.RedirectToErrorPage(w, handler.ErrorResponse{
			Title:   "Server Error",
			Message: "Couldn't retrieve account information. Please try again.",
		})
		return
	}

	if userData == nil || !h.passwordService.VerifyPassword(payload.Password, userData.Password) {
		data := map[string]any{
			"Form": payload,
			"Errors": map[string]string{
				"general": "Invalid email or password",
			},
			csrf.TemplateTag: csrf.TemplateField(r),
		}
		h.renderer.RenderTemplate(w, "login-form", data, "components/login-form.html")
		return
	}

	h.setCookie(w, userData.Id)
	w.Header().Set("HX-Redirect", "/")
}
