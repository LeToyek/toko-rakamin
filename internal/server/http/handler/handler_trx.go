package handler

import (
	"rakamin-final/internal/infrastructure/container"
	"rakamin-final/internal/pkg/controller"
	"rakamin-final/internal/pkg/repository"
	"rakamin-final/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
)

func RouteTrx(r fiber.Router, container *container.Container) {

	repo := repository.NewTrxRepository(container.MySqlDB)
	usecase := usecase.NewTrxUsecase(repo)
	controller := controller.NewTrxController(usecase)

	r.Group("/trx")
	r.Get("/", controller.GetAllTrxes)
	r.Get("/:id", controller.GetTrxByID)
	r.Post("/", controller.CreateTrx)
	r.Put("/:id", controller.UpdateTrx)
	r.Delete("/:id", controller.DeleteTrx)

}
