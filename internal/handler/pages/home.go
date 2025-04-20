package pages

import (
	"net/http"
)

func (p *Pages) Home(w http.ResponseWriter, r *http.Request) {
	p.renderer.Render(w, map[string]any{}, "layout.html", "pages/home.html")
}
