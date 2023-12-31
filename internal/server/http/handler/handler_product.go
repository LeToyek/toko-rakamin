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
	productApi := r.Group("/product").Use(middleware.DeserializeUser)

	repoProduct := repo.NewProductRepository(containerConf.MySqlDB)
	repoLog := repo.NewLogProductRepository(containerConf.MySqlDB)
	usecase := usecase.NewProductUsecase(repoProduct, repoLog)
	controller := controller.NewProductController(usecase)

	productApi.Get("/", controller.GetAllProducts)
	productApi.Post("/create", controller.CreateProduct)
	productApi.Get("/:id", controller.GetProductByID)
	productApi.Put("/edit/:id", controller.UpdateProduct)
	productApi.Delete("/delete/:id", controller.DeleteProduct)
}
