package handler

import (
	"rakamin-final/internal/infrastructure/container"
	"rakamin-final/internal/pkg/controller"
	"rakamin-final/internal/pkg/middleware"
	"rakamin-final/internal/pkg/repository"
	"rakamin-final/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
)

func RouteProductPhoto(r fiber.Router, container *container.Container) {
	repository := repository.NewProductPhotoRepository(container.MySqlDB)
	usecase := usecase.NewProductPhotoUsecase(repository)
	controller := controller.NewProductPhotoController(usecase)

	productPhotoApi := r.Group("/product-photo").Use(middleware.DeserializeUser)

	productPhotoApi.Get("/", controller.GetAllProductPhotos)
	productPhotoApi.Post("/create/:id", controller.CreateProductPhoto)
	productPhotoApi.Get("/:id", controller.GetProductPhotoByID)
	productPhotoApi.Put("/edit/:id", controller.UpdateProductPhoto)
	productPhotoApi.Delete("/delete/:id", controller.DeleteProductPhoto)

}
