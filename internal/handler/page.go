package handler

import (
	"net/http"
	"path/filepath"
	"text/template"
)

func RenderPage(w http.ResponseWriter, templ string, data any) {
	layoutPath := filepath.Join("templates", "layout.html")
	pagePath := filepath.Join("templates", "pages", templ+".html")

	tmpl, err := template.ParseFiles(layoutPath, pagePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func HomePage(w http.ResponseWriter, r *http.Request) {
	RenderPage(w, "home", nil)
}
