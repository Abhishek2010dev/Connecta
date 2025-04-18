package renderer

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

type Renderer interface {
	Render(w http.ResponseWriter, data any, templates ...string)
}

type templateRenderer struct {
	cache   map[string]*template.Template
	mu      sync.RWMutex
	dirName string
}

func New(dirName string) Renderer {
	return &templateRenderer{
		cache:   make(map[string]*template.Template),
		dirName: dirName,
	}
}

func (t *templateRenderer) Render(w http.ResponseWriter, data any, templates ...string) {
	key := filepath.Join(templates...)

	t.mu.RLock()
	tmpl, found := t.cache[key]
	t.mu.RUnlock()

	if !found {
		paths := make([]string, len(templates))
		for i, templatePath := range templates {
			paths[i] = filepath.Join(t.dirName, templatePath)
		}

		parsedTmpl, err := template.ParseFiles(paths...)
		if err != nil {
			log.Printf("template parsing error: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		t.mu.Lock()
		t.cache[key] = parsedTmpl
		t.mu.Unlock()

		tmpl = parsedTmpl
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Printf("template execution error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
