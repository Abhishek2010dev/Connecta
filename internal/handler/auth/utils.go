package auth

import (
	"log"
	"net/http"
	"time"

	"github.com/Abhishek2010dev/Connecta/internal/handler"
)

func (a *AuthHandler) setCookie(w http.ResponseWriter, userID int64) {
	token, err := a.sessionService.GenerateToken(userID)
	if err != nil {
		log.Println(err)
		handler.RedirectToErrorPage(w, handler.ErrorResponse{
			Title:   "Server Error",
			Message: "Couldn't create session. Try again.",
		})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     h.tokenKey,
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().Add(30 * 24 * time.Hour),
	})
}
