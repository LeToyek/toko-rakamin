package handler

import (
	"rakamin-final/internal/infrastructure/container"
	"rakamin-final/internal/pkg/controller"
	"rakamin-final/internal/pkg/middleware"
	"rakamin-final/internal/pkg/repository"
	"rakamin-final/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
)

func RouteCategory(r fiber.Router, container *container.Container) {

	repo := repository.NewCategoryRepository(container.MySqlDB)
	usecase := usecase.NewCategoryUsecase(repo)
	controller := controller.NewCategoryController(usecase)

	categoryApi := r.Group("/category")
	categoryApi.Get("/", middleware.DeserializeUser, controller.GetAllCategories)
	categoryApi.Get("/:id", middleware.DeserializeUser, controller.GetCategoryByID)
	categoryApi.Post("/", middleware.DeserializeUser, controller.CreateCategory)
	categoryApi.Put("/:id", middleware.DeserializeUser, controller.UpdateCategory)
	categoryApi.Delete("/:id", middleware.DeserializeUser, controller.DeleteCategory)
}
