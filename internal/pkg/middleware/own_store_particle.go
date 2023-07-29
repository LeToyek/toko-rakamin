package middleware

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CheckOwnStoreParticle(c *fiber.Ctx, storeParticleID int64) error {

	storeID := c.Locals("store_id").(string)
	id, err := strconv.ParseInt(storeID, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if id != storeParticleID {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	return nil
}
