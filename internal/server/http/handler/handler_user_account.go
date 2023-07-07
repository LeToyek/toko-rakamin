package handler

import (
	"rakamin-final/internal/infrastructure/container"
	controller "rakamin-final/internal/pkg/controller"
	repo "rakamin-final/internal/pkg/repository"
	usecase "rakamin-final/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
)

func RouteUserAccount(r fiber.Router, containerConf *container.Container) {
	userApi := r.Group("/user")

	repo := repo.NewUserRepository(containerConf.MySqlDB)
	usecase := usecase.NewUsersUsecase(repo)
	controller := controller.NewUsersController(usecase)

	userApi.Get("/", controller.GetUsers)
	userApi.Post("/register", controller.RegisterUser)
	userApi.Post("/login", controller.LoginUser)
	userApi.Put("/edit/:id", controller.UpdateUser)
	userApi.Delete("/delete/:id", controller.DeleteUser)
	userApi.Get("/logout", controller.LogoutUser)
}
