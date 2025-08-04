package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter() http.Handler {

	r := chi.NewRouter()
	addFiles(r)
	addHomeRoutes(r)

	return r
}
