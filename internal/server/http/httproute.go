package http

import (
	route "rakamin-final/internal/server/http/handler"

	"rakamin-final/internal/infrastructure/container"

	"github.com/gofiber/fiber/v2"
)

func HTTPRouteInit(r *fiber.App, containerConf *container.Container) {
	api := r.Group("/api/v1") // /api

	route.RouteCategory(api, containerConf)
	route.RouteStore(api, containerConf)
	route.RouteUserAccount(api, containerConf)
	route.RouteProduct(api, containerConf)
	route.RouteAddress(api, containerConf)
	route.RouteTrx(api, containerConf)
	route.RouteProductPhoto(api, containerConf)
}
