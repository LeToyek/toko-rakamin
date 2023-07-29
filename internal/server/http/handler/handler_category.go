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
	categoryApi.Use(middleware.DeserializeUser)

	categoryApi.Get("/", controller.GetAllCategories)
	categoryApi.Get("/:id", controller.GetCategoryByID)
	categoryApi.Post("/create", middleware.CheckAdmin, controller.CreateCategory)
	categoryApi.Put("/edit/:id", middleware.CheckAdmin, controller.UpdateCategory)
	categoryApi.Delete("/delete/:id", middleware.CheckAdmin, controller.DeleteCategory)
}
