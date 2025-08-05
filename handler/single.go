package handler

import (
	"golang_tailwinds/internal/render"
	"net/http"
)

func Single(w http.ResponseWriter, r *http.Request) {

	// currentURL := r.URL.Path
	var data = map[string]interface{}{
		"title":   "Single Page",
		"content": "hi single",
	}

	render.RenderTemplate(w, "home", data)
}
