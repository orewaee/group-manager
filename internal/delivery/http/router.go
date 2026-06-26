package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRouter(handler *Handler) http.Handler {
	r := chi.NewRouter()

	r.Use(LoggerMiddleware)
	r.Use(middleware.Recoverer)

	r.Route("/v1", func(r chi.Router) {
		r.Post("/people", handler.CreatePerson)
		r.Put("/people/{id}", handler.UpdatePerson)
		r.Delete("/people/{id}", handler.DeletePerson)
		r.Post("/groups", handler.CreateGroup)
		r.Put("/groups/{id}", handler.UpdateGroup)
		r.Delete("/groups/{id}", handler.DeleteGroup)
		r.Get("/groups", handler.ListGroups)
		r.Get("/groups/{id}/members", handler.GetGroupMembers)
	})

	return r
}
