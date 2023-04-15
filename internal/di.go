package di

import (
	"fmt"

	http_handlers "github.com/io-m/clean/internal/user/handlers/http"
	"github.com/io-m/clean/internal/user/repository"
	"github.com/io-m/clean/internal/user/router"
	user "github.com/io-m/clean/internal/user/services"
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
	userRepository := repository.NewInMemoryUserRepository()
	userService := user.NewUserService(userRepository)
	chiHandler := http_handlers.NewChiHandler(userService)
	chiRouter := router.NewChiRouter(chiHandler)
	
	chiRouter.Handle(fmt.Sprintf(":%d", app.port))
}
