package handler

import (
	"rakamin-final/internal/infrastructure/container"
	controller "rakamin-final/internal/pkg/controller"
	"rakamin-final/internal/pkg/middleware"
	repo "rakamin-final/internal/pkg/repository"
	usecase "rakamin-final/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
)

func RouteUserAccount(r fiber.Router, containerConf *container.Container) {

	userRepo := repo.NewUserRepository(containerConf.MySqlDB)
	storeRepo := repo.NewStoreRepository(containerConf.MySqlDB)
	usecase := usecase.NewUsersUsecase(userRepo, storeRepo)
	controller := controller.NewUsersController(usecase)

	userApi := r.Group("/user")

	userApi.Get("/", middleware.CheckAdmin, controller.GetUsers)
	userApi.Post("/register", controller.RegisterUser)
	userApi.Post("/login", controller.LoginUser)
	userApi.Put("/edit/:id", middleware.DeserializeUser, middleware.CheckOwnUser, controller.UpdateUser)
	userApi.Delete("/delete/:id", controller.DeleteUser)
	userApi.Get("/logout", controller.LogoutUser)

}
