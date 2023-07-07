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

	r.Group("/product-photo")
	r.Get("/", middleware.DeserializeUser, controller.GetAllProductPhotos)
	r.Post("/create", middleware.DeserializeUser, controller.CreateProductPhoto)
	r.Get("/:id", middleware.DeserializeUser, controller.GetProductPhotoByID)
	r.Put("/edit/:id", middleware.DeserializeUser, controller.UpdateProductPhoto)
	r.Delete("/delete/:id", middleware.DeserializeUser, controller.DeleteProductPhoto)

}
