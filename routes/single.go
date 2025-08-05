package routes

import (
	"golang_tailwinds/handler"

	"github.com/go-chi/chi/v5"
)


func addSingle(mux *chi.Mux) {
	mux.Get("/single",handler.Single)

}
