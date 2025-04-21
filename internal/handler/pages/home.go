package pages

import (
	"net/http"
)

func (p *Pages) Home(w http.ResponseWriter, r *http.Request) {
	payload := GetAuthPayload(r)
	data := map[string]any{
		"Username": payload.Username,
	}
	p.renderer.Render(w, data, "layout.html", "pages/home.html", "components/sidebar.html")
}
