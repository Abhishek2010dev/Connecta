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
	err := r.ParseForm()
	if err != nil {
		log.Printf("Failed to process request: %s", err)
		RedirectToErrorPage(w, ErrorResponse{
			Title:   "Something went wrong",
			Message: "We couldn't process your request. Please try again.",
		})
		return nil
	}

	var payload T
	if err := decoder.Decode(&payload, r.PostForm); err != nil {
		log.Printf("Invalid submission: %s", err)
		RedirectToErrorPage(w, ErrorResponse{
			Title:   "Invalid submission",
			Message: "There was an issue with the information you entered. Please review and try again.",
		})
		return nil
	}

	return &payload
}
