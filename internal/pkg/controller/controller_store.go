package controller

import (
	"fmt"
	"rakamin-final/internal/pkg/dto"
	"rakamin-final/internal/pkg/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type StoreController interface {
	CreateStore(ctx *fiber.Ctx) error
	GetAllStores(ctx *fiber.Ctx) error
	GetStoreByID(ctx *fiber.Ctx) error
	UpdateStore(ctx *fiber.Ctx) error
	DeleteStore(ctx *fiber.Ctx) error
}

type storeControllerImpl struct {
	usecase usecase.StoreUsecase
}

func NewStoreController(usecase usecase.StoreUsecase) *storeControllerImpl {
	return &storeControllerImpl{
		usecase: usecase,
	}
}

func (u *storeControllerImpl) CreateStore(ctx *fiber.Ctx) error {
	c := ctx.Context()

	var store dto.StoreRequest

	if err := ctx.BodyParser(&store); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	res, err := u.usecase.CreateStore(c, store)

	if err != nil {
		return ctx.Status(err.Code).JSON(fiber.Map{
			"message": err.Message.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
		"data":    res,
		"store":   store,
	})

}

func (u *storeControllerImpl) GetAllStores(ctx *fiber.Ctx) error {
	c := ctx.Context()

	res, err := u.usecase.GetAllStores(c, dto.StoreFilter{})

	fmt.Println(`workk`)
	if err != nil {
		fmt.Println(err)
		return ctx.Status(err.Code).JSON(fiber.Map{
			"location": "controller.GetAllStores",
			"message":  err.Message.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    res,
	})
}

func (u *storeControllerImpl) GetStoreByID(ctx *fiber.Ctx) error {
	c := ctx.Context()

	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	res, errUsecase := u.usecase.GetStoreByID(c, int64(id))

	if errUsecase != nil {
		return ctx.Status(errUsecase.Code).JSON(fiber.Map{
			"message": errUsecase.Message.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    res,
	})
}

func (u *storeControllerImpl) UpdateStore(ctx *fiber.Ctx) error {
	c := ctx.Context()

	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var store dto.StoreRequest

	if err := ctx.BodyParser(&store); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	res, errUsecase := u.usecase.UpdateStore(c, int64(id), store)

	if errUsecase != nil {
		return ctx.Status(errUsecase.Code).JSON(fiber.Map{
			"message": errUsecase.Message.Error(),
		})

	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    res,
		"store":   store,
	})

}

func (u *storeControllerImpl) DeleteStore(ctx *fiber.Ctx) error {
	c := ctx.Context()

	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	errUsecase := u.usecase.DeleteStore(c, int64(id))

	if errUsecase != nil {
		return ctx.Status(errUsecase.Code).JSON(fiber.Map{
			"message": errUsecase.Message.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})
}
