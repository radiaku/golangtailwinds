package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func addFiles(mux *chi.Mux) {

	FileServer(mux, "/files", http.Dir("files"))
}


func FileServer(r chi.Router, path string, root http.FileSystem) {
	fs := http.StripPrefix(path, http.FileServer(root))

	r.Get(path+"/*", func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	})
}
