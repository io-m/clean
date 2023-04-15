package router

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/io-m/clean/internal/core"
)

type handlers struct{
	authHandler core.IAuthHttpHandler
	userHandler core.IUserHttpHandler
}


type chiRouter struct {
	router  chi.Router
	handler handlers
}

func NewChiRouter(hs []any) IHttpRouter {
	var ah core.IAuthHttpHandler
	var uh core.IUserHttpHandler
	// making sure to assign proper handlers by checking type of them
	for _, h := range hs {
		a, okAuth := h.(core.IAuthHttpHandler) 
		u, okUser := h.(core.IUserHttpHandler) 
		if okAuth {
			ah = a
		}
		if okUser {
			uh = u
		}
	}
	return &chiRouter{
		router:  chi.NewRouter(),
		handler: handlers{
			authHandler: ah,
			userHandler: uh,
		},
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
		r.Get("/{id}", c.handler.userHandler.FindOne)
	})
	c.router.Route("/auth", func(r chi.Router) {
		r.Get("/{token}", c.handler.authHandler.Verify)
	})
}