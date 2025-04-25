package response

import (
	"github.com/gofiber/fiber/v2"
)

func FiberResponse(c *fiber.Ctx, code int, msg string, data interface{}) error {
	return c.Status(code).JSON(fiber.Map{
		"message": msg,
		"data":    data,
	})
}
