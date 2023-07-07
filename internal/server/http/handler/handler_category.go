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

	r.Group("/category")
	r.Get("/", middleware.DeserializeUser, controller.GetAllCategories)
	r.Get("/:id", middleware.DeserializeUser, controller.GetCategoryByID)
	r.Post("/", middleware.DeserializeUser, controller.CreateCategory)
	r.Put("/:id", middleware.DeserializeUser, controller.UpdateCategory)
	r.Delete("/:id", middleware.DeserializeUser, controller.DeleteCategory)
}
