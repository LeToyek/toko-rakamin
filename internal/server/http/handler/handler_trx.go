package handler

import (
	"rakamin-final/internal/infrastructure/container"
	"rakamin-final/internal/pkg/controller"
	"rakamin-final/internal/pkg/middleware"
	"rakamin-final/internal/pkg/repository"
	"rakamin-final/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
)

func RouteTrx(r fiber.Router, container *container.Container) {

	repo := repository.NewTrxRepository(container.MySqlDB)
	repoAddress := repository.NewAddressRepository(container.MySqlDB)
	repoProduct := repository.NewProductRepository(container.MySqlDB)
	repoDetailTrx := repository.NewDetailTrxRepository(container.MySqlDB)
	usecase := usecase.NewTrxUsecase(repo, repoDetailTrx, repoProduct, repoAddress)
	controller := controller.NewTrxController(usecase)

	trxApi := r.Group("/trx").Use(middleware.DeserializeUser)
	trxApi.Get("/", controller.GetAllTrxes)
	trxApi.Get("/:id", controller.GetTrxByID)
	trxApi.Post("/create", controller.CreateTrx)
	// trxApi.Put("/edit/:id", controller.UpdateTrx)
	trxApi.Delete("/delete/:id", controller.DeleteTrx)

}
