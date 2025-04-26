package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ntp7758/task-management/internal/auth/handler"
	"github.com/ntp7758/task-management/internal/auth/repository"
	"github.com/ntp7758/task-management/internal/auth/routes"
	"github.com/ntp7758/task-management/internal/auth/service"
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

	authService := service.NewAuthService(authRepo)

	authHandler := handler.NewAuthHandler(authService)

	authRoute := routes.NewAuthRoute(authHandler)

	authRoute.Install(app)

	app.Listen(fmt.Sprintf(":%s", port))
}
