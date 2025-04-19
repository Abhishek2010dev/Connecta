package middleware

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"

	"github.com/Abhishek2010dev/Connecta/internal/handler"
	"github.com/Abhishek2010dev/Connecta/internal/service"
)

type Auth struct {
	sessionService service.Session
}

func NewAuth(db *sql.DB) *Auth {
	return &Auth{service.NewSession(db)}
}

var AuthPayloadKey = "auth-payload"

func (a *Auth) RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("authToken")
		if err != nil {
			http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
			return
		}

		_, payload, err := a.sessionService.ValidateToken(cookie.Value)
		if err != nil {
			if errors.Is(err, service.ErrSessionExpired) {
				http.SetCookie(w, &http.Cookie{
					Name:     "authToken",
					Value:    "",
					Path:     "/",
					MaxAge:   -1,
					HttpOnly: true,
				})
				http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
				return
			}

			log.Println("Token validation failed:", err)

			handler.RedirectToErrorPage(w, handler.ErrorResponse{
				Title:   "Session Error",
				Message: "We couldn't verify your session. Please log in again.",
			})
			return
		}

		ctx := context.WithValue(r.Context(), AuthPayloadKey, *payload)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
