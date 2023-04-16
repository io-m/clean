package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/io-m/clean/internal/auth"
	"github.com/io-m/clean/internal/user"
	"github.com/io-m/clean/internal/user/db"
	"github.com/io-m/clean/internal/user/handlers"
)

type appSetup struct {
	port int
	mux  *chi.Mux
}

func AppSetup(port int) *appSetup {
	return &appSetup{
		port: port,
		mux:  chi.NewMux(),
	}
}

func (app *appSetup) GogoBaby() {
	/* Repositories */
	userRepository := db.NewInMemoryUserRepository()

	/* Services */
	userService := user.NewUserService(userRepository)
	authService := auth.NewAuthService(userService)

	/* Routes & Handlers */
	go handlers.NewUserHandler(app.mux, userService).HandleRoutes()
	go auth.NewAuthHandler(app.mux, authService).HandleRoutes()

	/* Server listening */
	log.Println("Server is listening on port -> ", app.port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", app.port), app.mux); err != nil {
		log.Fatal("Server is down")
		// gracefully shut down the server
	}
}
