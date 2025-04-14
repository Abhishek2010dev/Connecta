package handler

import (
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

var (
	templateCache = make(map[string]*template.Template)
	cacheMutex    sync.RWMutex
)

func render(w http.ResponseWriter, data any, templates ...string) {
	key := filepath.Join(templates...)
	cacheMutex.RLock()
	tmpl, found := templateCache[key]
	cacheMutex.RUnlock()

	if !found {
		paths := make([]string, len(templates))
		for i, t := range templates {
			paths[i] = filepath.Join("templates", t)
		}

		var err error
		tmpl, err = template.ParseFiles(paths...)
		if err != nil {
			log.Printf("template parsing error: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		cacheMutex.Lock()
		templateCache[key] = tmpl
		cacheMutex.Unlock()
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		log.Printf("template execution error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	render(w, nil, "layout.html", "pages/home.html")
}
