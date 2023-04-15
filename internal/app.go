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

	userRepository := db.NewInMemoryUserRepository()
	userService := user.NewUserService(userRepository)
	userHandler := user.NewUserHandler((userService))
	authService := auth.NewAuthService(userService)
	authHandler := auth.NewAuthHandler(authService)
	chiRouter:= router.NewChiRouter(authHandler, userHandler)

	chiRouter.Handle(fmt.Sprintf(":%d", app.port))
}
