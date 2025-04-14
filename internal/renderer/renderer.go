package renderer

import (
	"net/http"
	"sync"
	"text/template"
)

type Renderer interface {
	Render(w http.ResponseWriter, data any, templates ...string)
}

type templateRenderer struct {
	cache   map[string]*template.Template
	mu      sync.RWMutex
	dirName string
}

func New() Renderer {
	return &templateRenderer{}
}

func (t *templateRenderer) Render(w http.ResponseWriter, data any, templates ...string) {
	panic("not implemented") // TODO: Implement
}
