package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ntp7758/task-management/internal/auth/model"
	"github.com/ntp7758/task-management/internal/auth/service"
	"github.com/ntp7758/task-management/pkg/response"
)

type AuthHandler interface {
	Signup(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}

type authHandler struct {
	authService service.AuthService
}

func NewAuthHandler(s service.AuthService) AuthHandler {
	return &authHandler{s}
}

func (h *authHandler) Signup(c *fiber.Ctx) error {
	var req model.SignupRequest
	err := c.BodyParser(&req)
	if err != nil {
		return response.FiberResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	if req.Password != req.ConfirmPassword {
		return response.FiberResponse(c, fiber.StatusBadRequest, "password and confirm is not invalid", nil)
	}

	err = h.authService.Signup(req.Username, req.Password)
	if err != nil {
		return response.FiberResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return response.FiberResponse(c, fiber.StatusCreated, "create success", nil)
}

func (h *authHandler) Login(c *fiber.Ctx) error {
	var req model.LoginRequest
	err := c.BodyParser(&req)
	if err != nil {
		return response.FiberResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	}
	token, refreshToten, err := h.authService.Login(req.Username, req.Password)
	if err != nil {
		return response.FiberResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return response.FiberResponse(c, fiber.StatusOK, "success", map[string]string{
		"token":        token,
		"refreshToten": refreshToten,
	})
}
