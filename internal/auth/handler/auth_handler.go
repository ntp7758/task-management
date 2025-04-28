package handler

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/ntp7758/task-management/internal/auth/model"
	"github.com/ntp7758/task-management/internal/auth/service"
	userService "github.com/ntp7758/task-management/internal/user/service"
	"github.com/ntp7758/task-management/pkg/response"
)

type AuthHandler interface {
	Signup(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	CheckToken(c *fiber.Ctx) error
}

type authHandler struct {
	authService service.AuthService
	userService userService.UserService
}

func NewAuthHandler(authService service.AuthService, userService userService.UserService) AuthHandler {
	return &authHandler{authService, userService}
}

func (h *authHandler) Signup(c *fiber.Ctx) error {
	var req model.SignupRequest
	err := c.BodyParser(&req)
	if err != nil {
		return response.FiberResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	if strings.TrimSpace(req.Username) == "" || strings.TrimSpace(req.Password) == "" || strings.TrimSpace(req.ConfirmPassword) == "" {
		return response.FiberResponse(c, fiber.StatusBadRequest, "have empty input", nil)
	}

	if req.Password != req.ConfirmPassword {
		return response.FiberResponse(c, fiber.StatusBadRequest, "password and confirm is not invalid", nil)
	}

	authId, err := h.authService.Signup(req.Username, req.Password)
	if err != nil {
		return response.FiberResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	err = h.userService.Create(authId)
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

	authId, err := h.authService.Login(req.Username, req.Password)
	if err != nil {
		return response.FiberResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	user, err := h.userService.GetByAuthId(authId)
	if err != nil {
		return response.FiberResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	token, refreshToten, err := h.authService.CreateToken(user.ID.Hex())
	if err != nil {
		return response.FiberResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return response.FiberResponse(c, fiber.StatusOK, "success", map[string]string{
		"token":        token,
		"refreshToten": refreshToten,
	})
}

func (h *authHandler) CheckToken(c *fiber.Ctx) error {
	id, ok := c.Locals("id").(string)
	if !ok {
		return response.FiberResponse(c, fiber.StatusUnauthorized, "invalid token", nil)
	}

	return response.FiberResponse(c, fiber.StatusOK, "success", map[string]string{
		"id": id,
	})
}
