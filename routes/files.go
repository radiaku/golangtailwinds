package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func addFiles(mux *chi.Mux) {

	assets := http.FileServer(http.Dir("./assets/"))
	static := http.FileServer(http.Dir("./static/"))
	files := http.FileServer(http.Dir("./files/"))
	mux.Handle("/assets/*", http.StripPrefix("/assets", assets))
	mux.Handle("/static/*", http.StripPrefix("/static", static))
	mux.Handle("/files/*", http.StripPrefix("/files", files))

}
