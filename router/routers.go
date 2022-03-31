package router

import (
	"go-chi/API"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var r *chi.Mux

func init() {
	r = chi.NewRouter()
	r.Use(middleware.Logger)
}

func Routers() *chi.Mux {
	r.Get("/", API.GetBlog)
	r.Get("/list", API.GetListBlog)
	return r
}
