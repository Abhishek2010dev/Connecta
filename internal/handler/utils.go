package handler

import (
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/schema"
)

type ErrorResponse struct {
	Title   string
	Message string
}

func RedirectToErrorPage(w http.ResponseWriter, error ErrorResponse) {
	params := url.Values{}
	params.Add("title", error.Title)
	params.Add("message", error.Message)
	target := "/error?" + params.Encode()

	w.Header().Set("HX-Redirect", target)
}

func ParseAndDecodeForm[T any](w http.ResponseWriter, r *http.Request, decoder *schema.Decoder) *T {
	if err := r.ParseForm(); err != nil {
		log.Printf("Failed to parse form: %s", err)
		RedirectToErrorPage(w, ErrorResponse{
			Title:   "Request Error",
			Message: "Couldn't process your request. Please try again.",
		})
		return nil
	}

	var payload T
	if err := decoder.Decode(&payload, r.PostForm); err != nil {
		log.Printf("Failed to decode form: %s", err)
		RedirectToErrorPage(w, ErrorResponse{
			Title:   "Invalid Data",
			Message: "Something’s wrong with your input. Please check and try again.",
		})
		return nil
	}

	return &payload
}
