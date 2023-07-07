package http

import (
	route "rakamin-final/internal/server/http/handler"

	"rakamin-final/internal/infrastructure/container"

	"github.com/gofiber/fiber/v2"
)

func HTTPRouteInit(r *fiber.App, containerConf *container.Container) {
	api := r.Group("/api/v1") // /api

	route.RouteStore(api, containerConf)
	route.RouteUserAccount(api, containerConf)
	route.RouteProduct(api, containerConf)
	route.RouteAddress(api, containerConf)
}
