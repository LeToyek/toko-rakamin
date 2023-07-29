package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func CheckOwnUser(c *fiber.Ctx) error {

	userID := c.Locals("user_id").(string)

	paramID := c.Params("id")

	if userID != paramID {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	return c.Next()
}
