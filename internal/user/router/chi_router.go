package router

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/io-m/clean/internal/user/core"
)

type chiRouter struct {
	router  chi.Router
	handler core.IHttpHandler
}

func NewChiRouter(h core.IHttpHandler) IHttpRouter {
	return &chiRouter{
		router:  chi.NewRouter(),
		handler: h,
	}
}

func (c *chiRouter) Handle(port string) {
	log.Printf("Server is listening on port 127.0.0%s \n", port)
	go c.handleRoutes()
	if err := http.ListenAndServe(port, c.router); err != nil {
		log.Fatal("Server is down")
		// gracefully shut down the server
	}
}

func (c *chiRouter) handleRoutes() {
	c.router.Route("/user", func(r chi.Router) {
		// r.Get("/", c.handler.FindOne)
		r.Get("/{id}", c.handler.FindOne)
	})
	c.router.Route("/auth", func(r chi.Router) {
		// r.Get("/login", c.handler.FindOne)
	})
}