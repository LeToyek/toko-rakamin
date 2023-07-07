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

	trxApi := r.Group("/trx")
	trxApi.Get("/", controller.GetAllTrxes)
	trxApi.Get("/:id", controller.GetTrxByID)
	trxApi.Post("/", controller.CreateTrx)
	trxApi.Put("/:id", controller.UpdateTrx)
	trxApi.Delete("/:id", controller.DeleteTrx)

}
