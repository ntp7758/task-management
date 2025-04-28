package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ntp7758/task-management/internal/user/service"
	"github.com/ntp7758/task-management/pkg/response"
)

type UserHandler interface {
	GetUser(c *fiber.Ctx) error
}

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(s service.UserService) UserHandler {
	return &userHandler{s}
}

func (h *userHandler) GetUser(c *fiber.Ctx) error {
	id, ok := c.Locals("id").(string)
	if !ok {
		return response.FiberResponse(c, fiber.StatusUnauthorized, "invalid token", nil)
	}

	user, err := h.userService.GetByUserId(id)
	if err != nil {
		return response.FiberResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return response.FiberResponse(c, fiber.StatusOK, "success", map[string]interface{}{
		"user": user,
	})
}
