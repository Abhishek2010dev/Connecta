package renderer

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"sync"
)

type Renderer interface {
	Render(w http.ResponseWriter, data any, templates ...string)
	RenderTemplate(w http.ResponseWriter, name string, data any, templates ...string)
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

// loadTemplate caches and returns a parsed template
func (t *templateRenderer) loadTemplate(templates ...string) (*template.Template, error) {
	key := strings.Join(templates, "|")

	t.mu.RLock()
	tmpl, found := t.cache[key]
	t.mu.RUnlock()

	if found {
		return tmpl, nil
	}

	paths := make([]string, len(templates))
	for i, file := range templates {
		paths[i] = filepath.Join(t.dirName, file)
	}

	parsedTmpl, err := template.ParseFiles(paths...)
	if err != nil {
		return nil, err
	}

	t.mu.Lock()
	t.cache[key] = parsedTmpl
	t.mu.Unlock()

	return parsedTmpl, nil
}

func (t *templateRenderer) Render(w http.ResponseWriter, data any, templates ...string) {
	tmpl, err := t.loadTemplate(templates...)
	if err != nil {
		log.Printf("template load error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Printf("template execution error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (t *templateRenderer) RenderTemplate(w http.ResponseWriter, name string, data any, templates ...string) {
	tmpl, err := t.loadTemplate(templates...)
	if err != nil {
		log.Printf("template load error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := tmpl.ExecuteTemplate(w, name, data); err != nil {
		log.Printf("template execution error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

