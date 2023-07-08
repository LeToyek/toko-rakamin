package controller

import (
	"fmt"
	"rakamin-final/internal/pkg/dto"
	"rakamin-final/internal/pkg/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type AddressController interface {
	CreateAddress(ctx *fiber.Ctx) error
	GetAddress(ctx *fiber.Ctx) error
	UpdateAddress(ctx *fiber.Ctx) error
	DeleteAddress(ctx *fiber.Ctx) error
}

type addressControllerImpl struct {
	usecase usecase.AddressUsecase
}

func NewAddressController(usecase usecase.AddressUsecase) *addressControllerImpl {
	return &addressControllerImpl{
		usecase: usecase,
	}
}

func (u *addressControllerImpl) CreateAddress(ctx *fiber.Ctx) error {
	c := ctx.Context()

	var address dto.AddressRequest

	if err := ctx.BodyParser(&address); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	res, err := u.usecase.CreateAddress(c, address)

	if err != nil {
		return ctx.Status(err.Code).JSON(fiber.Map{
			"message": err.Message.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
		"id":      res.ID,
		"address": address,
	})

}

func (u *addressControllerImpl) GetAddress(ctx *fiber.Ctx) error {
	c := ctx.Context()

	id, _ := strconv.Atoi(ctx.Params("id"))

	res, err := u.usecase.GetAddressByID(c, int64(id))

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

func (u *addressControllerImpl) UpdateAddress(ctx *fiber.Ctx) error {
	c := ctx.Context()

	id, _ := strconv.Atoi(ctx.Params("id"))

	var address dto.AddressRequest

	if err := ctx.BodyParser(&address); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	_, errUC := u.usecase.UpdateAddress(c, int64(id), dto.AddressRequest{
		JudulAlamat:  address.JudulAlamat,
		NamaPenerima: address.NamaPenerima,
		NoTelp:       address.NoTelp,
		DetailAlamat: address.DetailAlamat,
	})

	if errUC != nil {
		return ctx.Status(errUC.Code).JSON(fiber.Map{
			"message": errUC.Message.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"address": address,
	})
}

func (u *addressControllerImpl) DeleteAddress(ctx *fiber.Ctx) error {
	c := ctx.Context()

	id, _ := strconv.Atoi(ctx.Params("id"))

	err := u.usecase.DeleteAddress(c, int64(id))

	if err != nil {
		return ctx.Status(err.Code).JSON(fiber.Map{
			"message": err.Message.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("success delete address with id %d", id),
	})
}
