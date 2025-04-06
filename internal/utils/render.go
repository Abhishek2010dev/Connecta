package utils

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func RenderPage(w http.ResponseWriter, page string, data interface{}) {
	files := []string{
		filepath.Join("templates", "layout.html"),
		filepath.Join("templates/pages", page+".html"),
	}
	tmpl := template.Must(template.ParseFiles(files...))
	tmpl.ExecuteTemplate(w, "layout", data)
}
