package handler

import (
	"rakamin-final/internal/infrastructure/container"
	"rakamin-final/internal/pkg/middleware"

	controller "rakamin-final/internal/pkg/controller"
	repo "rakamin-final/internal/pkg/repository"
	usecase "rakamin-final/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
)

func RouteAddress(r fiber.Router, container *container.Container) {
	addressApi := r.Group("/address")

	repo := repo.NewAddressRepository(container.MySqlDB)
	usecase := usecase.NewAddressUsecase(repo)
	controller := controller.NewAddressController(usecase)

	addressApi.Delete("/delete/:id", middleware.DeserializeUser, controller.DeleteAddress)
	addressApi.Post("/create", middleware.DeserializeUser, controller.CreateAddress)
	addressApi.Get("/:id", middleware.DeserializeUser, controller.GetAddress)
	addressApi.Put("/edit/:id", middleware.DeserializeUser, controller.UpdateAddress)
}
