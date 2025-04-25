package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ntp7758/task-management/internal/auth/handler"
)

type Routes interface {
	Install(app *fiber.App)
}

type authRoutes struct {
	authHandler handler.AuthHandler
}

func NewAuthRoute(authHandler handler.AuthHandler) Routes {
	return &authRoutes{authHandler: authHandler}
}

func (r *authRoutes) Install(app *fiber.App) {
	prefix := "/auth"

	app.Post(prefix+"/sign-up", r.authHandler.Signup)
	app.Post(prefix+"/login", r.authHandler.Login)
}
