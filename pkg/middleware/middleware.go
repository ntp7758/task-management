package middleware

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/ntp7758/task-management/pkg/response"
	"github.com/ntp7758/task-management/pkg/security"
)

func AuthMiddleware(c *fiber.Ctx) error {
	tokenStr := c.Get("Authorization")
	token := strings.TrimPrefix(tokenStr, "Bearer ")

	if token == "" {
		log.Println("authorization header is empty")
		return response.FiberResponse(c, fiber.StatusUnauthorized, "authorization header is empty", nil)
	}

	claims, err := security.ParseJWTToken(tokenStr)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "unauthorized",
		})
	}

	c.Locals("id", claims.UserID)

	return c.Next()
}
