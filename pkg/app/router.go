package app

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/hamologist/dice-roll/pkg/router"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Mount("/roll", router.RollRouter())

	return r
}
