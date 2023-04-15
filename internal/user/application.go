package user

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/io-m/clean/internal/user/core"
)

type chiRouter struct {
	router  *chi.Mux
	handler core.IHttpHandler
}

func NewChiRouter(h core.IHttpHandler) core.IHttpRouter {
	return &chiRouter{
		router:  chi.NewMux(),
		handler: h,
	}
}

func (cr *chiRouter) Serve(port string) error {
	log.Printf("Server is listening on port 127.0.0%s \n", port)
	return http.ListenAndServe(port, cr.router)
}

func (cr *chiRouter) HandleRoutes() {
	cr.router.Route("/user", func(r chi.Router) {
		r.Get("/", cr.handler.FindOne)
		r.Get("/{id}", cr.handler.FindOne)
	})
	cr.router.Route("/auth", func(r chi.Router) {
		r.Get("/login", cr.handler.FindOne)
	})
}
