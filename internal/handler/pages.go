package handler

import (
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

func render(w http.ResponseWriter, data any, templates ...string) {
	paths := make([]string, len(templates))
	for i, t := range templates {
		paths[i] = filepath.Join("templates", t)
	}

	tmpl, err := template.ParseFiles(paths...)
	if err != nil {
		log.Printf("template parsing error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("template execution error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	render(w, nil, "layout.html", "pages/home.html")
}
