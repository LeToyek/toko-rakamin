package controller

import (
	"fmt"
	userDTO "rakamin-final/internal/pkg/dto"
	"rakamin-final/internal/pkg/usecase"
	"time"

	"github.com/gofiber/fiber/v2"
)

type UsersController interface {
	LoginUser(ctx *fiber.Ctx) error
	RegisterUser(ctx *fiber.Ctx) error
	GetUsers(ctx *fiber.Ctx) error
	UpdateUser(ctx *fiber.Ctx) error
	DeleteUser(ctx *fiber.Ctx) error
	LogoutUser(ctx *fiber.Ctx) error
}

type usersControllerImpl struct {
	usecase usecase.UsersUsecase
}

func NewUsersController(usecase usecase.UsersUsecase) *usersControllerImpl {
	return &usersControllerImpl{
		usecase: usecase,
	}
}

func (u *usersControllerImpl) RegisterUser(ctx *fiber.Ctx) error {
	c := ctx.Context()

	var user userDTO.UserRegister

	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	res, err := u.usecase.CreateUser(c, user)

	if err != nil {
		return ctx.Status(err.Code).JSON(fiber.Map{
			"message": err.Message.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
		"data":    res,
		"user":    user,
	})

}

func (u *usersControllerImpl) LoginUser(ctx *fiber.Ctx) error {
	c := ctx.Context()

	var user userDTO.UserLogin

	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	res, err, token := u.usecase.GetCredentialUserLogin(c, userDTO.UserLogin{
		Email:    user.Email,
		Password: user.Password,
	})

	if err != nil {
		return ctx.Status(err.Code).JSON(fiber.Map{
			"message": err.Message.Error(),
		})
	}

	ctx.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(time.Minute * 60),
	})

	ctx.Response().Header.VisitAllCookie(func(key, value []byte) {
		fmt.Println("res cookieKey", string(key), "value", string(value))
	})

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    res,
	})

}

func (u *usersControllerImpl) GetUsers(ctx *fiber.Ctx) error {
	c := ctx.Context()

	params := new(userDTO.UserFilter)

	if err := ctx.QueryParser(params); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	res, err := u.usecase.GetAllUsers(c, *params)

	if err != nil {
		return ctx.Status(err.Code).JSON(fiber.Map{
			"message": err.Message.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    res,
	})
}

func (u *usersControllerImpl) UpdateUser(ctx *fiber.Ctx) error {
	c := ctx.Context()

	EXAMPLE_PARAM := 1

	var user userDTO.UserRegister

	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	res, err := u.usecase.UpdateUser(c, int64(EXAMPLE_PARAM), userDTO.UserRegister{
		Nama:         user.Nama,
		NoTelp:       user.NoTelp,
		JenisKelamin: user.JenisKelamin,
		Tentang:      user.Tentang,
		Pekerjaan:    user.Pekerjaan,
		IdProvinsi:   user.IdProvinsi,
		IdKota:       user.IdKota,
	})

	if err != nil {
		return ctx.Status(err.Code).JSON(fiber.Map{
			"message": err.Message.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success to update user data",
		"data":    res,
	})
}

func (u *usersControllerImpl) DeleteUser(ctx *fiber.Ctx) error {
	c := ctx.Context()

	const EXAMPLE_PARAM = 1

	err := u.usecase.DeleteUser(c, EXAMPLE_PARAM)

	if err != nil {
		return ctx.Status(err.Code).JSON(fiber.Map{
			"message": err.Message.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data has been deleted",
	})
}

func (u *usersControllerImpl) LogoutUser(ctx *fiber.Ctx) error {

	ctx.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Now().Add(-3 * time.Second),
	})

	ctx.ClearCookie("token")

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Logout success",
		"cookie":  ctx.Cookies("token"),
	})
}
