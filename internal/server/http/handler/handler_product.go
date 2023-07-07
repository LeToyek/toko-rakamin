package handler

import (
	"rakamin-final/internal/infrastructure/container"
	controller "rakamin-final/internal/pkg/controller"
	"rakamin-final/internal/pkg/middleware"
	repo "rakamin-final/internal/pkg/repository"
	usecase "rakamin-final/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
)

func RouteProduct(r fiber.Router, containerConf *container.Container) {
	productApi := r.Group("/product")

	repo := repo.NewProductRepository(containerConf.MySqlDB)
	usecase := usecase.NewProductUsecase(repo)
	controller := controller.NewProductController(usecase)

	productApi.Get("/", middleware.DeserializeUser, controller.GetAllProducts)
	productApi.Post("/create", middleware.DeserializeUser, controller.CreateProduct)
	productApi.Get("/:id", middleware.DeserializeUser, controller.GetProductByID)
	productApi.Put("/edit/:id", middleware.DeserializeUser, controller.UpdateProduct)
	productApi.Delete("/delete/:id", middleware.DeserializeUser, controller.DeleteProduct)
}
