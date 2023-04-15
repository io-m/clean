package main

import (
	"log"

	"github.com/io-m/clean/internal/user"
	http_handlers "github.com/io-m/clean/internal/user/handlers/http"
	"github.com/io-m/clean/internal/user/repository"
	service "github.com/io-m/clean/internal/user/services"
)

func main() {
	userRepository := repository.NewInMemoryUserRepository()
	userService := service.NewUserService(userRepository)
	chiHandler := http_handlers.NewChiHandler(userService)
	chiRouter := user.NewChiRouter(chiHandler)

	if err := chiRouter.Handle(":9090"); err != nil {
		log.Fatal("Server failed to start listening")
	}
}
