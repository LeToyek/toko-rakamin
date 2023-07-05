package handler

import (
	"rakamin-final/internal/infrastructure/container"

	"github.com/gofiber/fiber/v2"
)

func RouteStore(r fiber.Router, containerConf *container.Container) {
	storeAPI := r.Group("/store")
	storeAPI.Get("/", func(c *fiber.Ctx) error {
		c.SendString("Hello from store")
		return nil
	})
}
