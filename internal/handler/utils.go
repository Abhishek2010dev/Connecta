package handler

import (
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/schema"
)

type ErrorResponse struct {
	Title   string
	Message string
}

func redirectToErrorPage(w http.ResponseWriter, error ErrorResponse) {
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

func parseAndDecodeForm[T any](w http.ResponseWriter, r *http.Request, decoder *schema.Decoder) *T {
	err := r.ParseForm()
	if err != nil {
		log.Printf("Failed to process request: %s", err)
		redirectToErrorPage(w, ErrorResponse{
			Title:   "Something went wrong",
			Message: "We couldn't process your request. Please try again.",
		})
		return nil
	}

	var payload T
	if err := decoder.Decode(&payload, r.PostForm); err != nil {
		log.Printf("Invalid submission: %s", err)
		redirectToErrorPage(w, ErrorResponse{
			Title:   "Invalid submission",
			Message: "There was an issue with the information you entered. Please review and try again.",
		})
		return nil
	}

	return &payload
}
