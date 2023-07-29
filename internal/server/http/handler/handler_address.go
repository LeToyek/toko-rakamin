package handler

import (
	"rakamin-final/internal/infrastructure/container"

	controller "rakamin-final/internal/pkg/controller"
	"rakamin-final/internal/pkg/middleware"
	repo "rakamin-final/internal/pkg/repository"
	usecase "rakamin-final/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
)

func RouteAddress(r fiber.Router, container *container.Container) {
	addressApi := r.Group("/address")

	repo := repo.NewAddressRepository(container.MySqlDB)
	usecase := usecase.NewAddressUsecase(repo)
	controller := controller.NewAddressController(usecase)

	addressApi.Use(middleware.DeserializeUser)

	addressApi.Delete("/delete/:id", controller.DeleteAddress)
	addressApi.Post("/create", controller.CreateAddress)
	addressApi.Get("/:id", controller.GetAddress)
	addressApi.Put("/edit/:id", controller.UpdateAddress)
}
