package handler

import (
	"net/http"
	"net/url"
	"time"
)

type ErrorResposne struct {
	Title   string
	Message string
}

func redirectToErrorPage(w http.ResponseWriter, error ErrorResposne) {
	params := url.Values{}
	params.Add("title", error.Title)
	params.Add("message", error.Message)
	target := "/error?" + params.Encode()

	w.Header().Set("HX-Redirect", target)
}

func setCookie(w http.ResponseWriter, tokenName, token string) {
	http.SetCookie(w, &http.Cookie{
		Name:     tokenName,
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().Add(30 * 24 * time.Hour),
	})
}
