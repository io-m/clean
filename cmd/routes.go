package main

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/io-m/clean/internal/user/core"
)

func HandleRoutes(c any, h core.IHttpHandler) error {
	log.Print("HEREEE")
	router, ok := c.(chi.Router)
	if !ok {
		return core.ErrorHandlerNotSet
	}
	router.Route("/user", func(r chi.Router) {
		r.Get("/", h.FindOne)
		r.Get("/{id}", h.FindOne)
	})
	router.Route("/auth", func(r chi.Router) {
		r.Get("/login", h.FindOne)
	})
	return nil
}
