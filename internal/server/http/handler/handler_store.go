package handler

import (
	"rakamin-final/internal/infrastructure/container"
	"rakamin-final/internal/pkg/controller"
	"rakamin-final/internal/pkg/middleware"
	repo "rakamin-final/internal/pkg/repository"
	"rakamin-final/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
)

func RouteStore(r fiber.Router, containerConf *container.Container) {

	repo := repo.NewStoreRepository(containerConf.MySqlDB)
	usecase := usecase.NewStoreUsecase(repo)
	controller := controller.NewStoreController(usecase)

	storeAPI := r.Group("/store")

	storeAPI.Get("/", middleware.DeserializeUser, controller.GetAllStores)
	storeAPI.Get("/:id", middleware.DeserializeUser, controller.GetStoreByID)
	storeAPI.Post("/", middleware.DeserializeUser, controller.CreateStore)
	storeAPI.Put("/:id", middleware.DeserializeUser, controller.UpdateStore)
	storeAPI.Delete("/:id", middleware.DeserializeUser, controller.DeleteStore)
}
