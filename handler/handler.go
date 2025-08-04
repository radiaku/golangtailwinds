package handler

import (
	"golang_tailwinds/internal/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {

	// currentURL := r.URL.Path
	var data = map[string]interface{}{
		"title":   "Dashboard",
		"content": "hi folks",
	}

	render.RenderTemplate(w, "home", data)
}
