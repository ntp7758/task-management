package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ntp7758/task-management/internal/user/handler"
	"github.com/ntp7758/task-management/pkg/middleware"
)

type Routes interface {
	Install(app *fiber.App)
}

type userRoutes struct {
	userHandler handler.UserHandler
}

func NewUserRoute(userHandler handler.UserHandler) Routes {
	return &userRoutes{userHandler: userHandler}
}

func (r *userRoutes) Install(app *fiber.App) {
	prefix := "/user"

	app.Get(prefix+"/get-user", middleware.AuthMiddleware, r.userHandler.GetUser)
}
