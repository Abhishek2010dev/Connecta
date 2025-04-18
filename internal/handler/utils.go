package handler

import (
	"net/http"
	"net/url"
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
