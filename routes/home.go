package routes

import (
	"github.com/go-chi/chi/v5"
	"golang_tailwinds/handler"
)

func addHomeRoutes(mux *chi.Mux) {
	mux.Get("/",handler.Home)

}
