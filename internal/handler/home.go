package handler

import (
	"net/http"

	"github.com/Abhishek2010dev/Connecta/internal/renderer"
)

type HomeHandler struct {
	renderer renderer.Renderer
}

func NewHomeHandler(renderer renderer.Renderer) *HomeHandler {
	return &HomeHandler{renderer}
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.renderer.Render(w, nil, "layout.html", "pages/home.html")
}
