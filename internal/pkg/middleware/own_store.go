package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func CheckOwnStore(c *fiber.Ctx) error {

	storeID := c.Locals("store_id").(string)
	paramID := c.Params("id")

	if storeID != paramID {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	return c.Next()
}
