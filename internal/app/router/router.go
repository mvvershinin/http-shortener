package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/mvvershinin/http-shortener/internal/app/handler"
)

func GetRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Get(`/{uid}`, handler.GetHandler)
	router.Post(`/`, handler.PostHandler)
	router.NotFound(handler.BadRequestHandler)
	router.MethodNotAllowed(handler.BadRequestHandler)

	return router
}
