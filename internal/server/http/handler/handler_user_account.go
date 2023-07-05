package handler

import (
	"github.com/gofiber/fiber/v2"
)

func RouteUserAccount(r fiber.Router) {
	userApi := r.Group("/user")

	userApi.Get("/", func(c *fiber.Ctx) error {
		c.SendString("Hello from user")
		return nil
	})
}
