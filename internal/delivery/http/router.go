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
		r.Put("/people/:id", handler.UpdatePerson)
		// r.Delete("/people/:id", handler.DeletePerson)
	})

	return r
}
