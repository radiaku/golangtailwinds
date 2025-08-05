package routes

import (
	"github.com/go-chi/chi/v5"
	"golang_tailwinds/handler"
)

func addHome(mux *chi.Mux) {
	mux.Get("/",handler.Home)

}
