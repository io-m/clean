package app

import (
	"fmt"

	"github.com/io-m/clean/internal/auth"
	"github.com/io-m/clean/internal/db"
	"github.com/io-m/clean/internal/router"
	"github.com/io-m/clean/internal/user"
)

type appSetup struct {
	port int
}

func AppSetup(port int) *appSetup {
	return &appSetup{
		port: port,
	}
}

func(app *appSetup) GogoBaby() {

	/* Repositories */
	userRepository := db.NewInMemoryUserRepository()

	/* Services */
	userService := user.NewUserService(userRepository)
	authService := auth.NewAuthService(userService)
	
	/* Handlers */
	userHandler := user.NewUserHandler((userService))
	authHandler := auth.NewAuthHandler(authService)
	handlers := []any{authHandler, userHandler}
	/* Router */
	chiRouter := router.NewChiRouter(handlers)
	
	chiRouter.Handle(fmt.Sprintf(":%d", app.port))
}
