package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ntp7758/task-management/internal/auth/handler"
	"github.com/ntp7758/task-management/internal/auth/repository"
	"github.com/ntp7758/task-management/internal/auth/routes"
	"github.com/ntp7758/task-management/internal/auth/service"
	userHandler "github.com/ntp7758/task-management/internal/user/handler"
	userRepo "github.com/ntp7758/task-management/internal/user/repository"
	userRoute "github.com/ntp7758/task-management/internal/user/routes"
	userService "github.com/ntp7758/task-management/internal/user/service"
	"github.com/ntp7758/task-management/pkg/config"
	"github.com/ntp7758/task-management/pkg/db"
)

const (
	port string = "8080"
)

func main() {
	app := fiber.New()
	config.FiberConfig(app)

	client, err := db.NewMongoDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	authRepo, err := repository.NewAuthRepository(client)
	if err != nil {
		log.Fatal(err)
	}
	userRepo, err := userRepo.NewUserRepository(client)
	if err != nil {
		log.Fatal(err)
	}

	authService := service.NewAuthService(authRepo)
	userService := userService.NewUserService(userRepo)

	authHandler := handler.NewAuthHandler(authService, userService)
	userHandler := userHandler.NewUserHandler(userService)

	authRoute := routes.NewAuthRoute(authHandler)
	userRoute := userRoute.NewUserRoute(userHandler)

	authRoute.Install(app)
	userRoute.Install(app)

	app.Listen(fmt.Sprintf(":%s", port))
}
