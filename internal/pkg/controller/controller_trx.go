package controller

import (
	"fmt"
	"rakamin-final/internal/pkg/dto"
	"rakamin-final/internal/pkg/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type TrxController interface {
	CreateTrx(ctx *fiber.Ctx) error
	GetAllTrxes(ctx *fiber.Ctx) error
	GetTrxByID(ctx *fiber.Ctx) error
	// UpdateTrx(ctx *fiber.Ctx) error
	DeleteTrx(ctx *fiber.Ctx) error
}

type trxControllerImpl struct {
	usecase usecase.TrxUsecase
}

func NewTrxController(usecase usecase.TrxUsecase) *trxControllerImpl {
	return &trxControllerImpl{usecase}
}

func (u *trxControllerImpl) CreateTrx(ctx *fiber.Ctx) error {
	var trxData dto.TrxRequest
	if err := ctx.BodyParser(&trxData); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	resUserID := ctx.Locals("user_id").(string)
	fmt.Println(resUserID)
	userID, err := strconv.Atoi(resUserID)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	res, resErr := u.usecase.CreateTrx(ctx.Context(), trxData, int64(userID))

	if resErr != nil {
		return ctx.Status(resErr.Code).JSON(fiber.Map{
			"message": resErr.Message.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
		"data":    res,
	})
}

func (u *trxControllerImpl) GetAllTrxes(ctx *fiber.Ctx) error {
	c := ctx.Context()

	res, err := u.usecase.GetAllTrxes(c, dto.FilterTrx{
		Limit:  ctx.QueryInt("limit", 10),
		Offset: ctx.QueryInt("offset", 0),
	})

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

func (u *trxControllerImpl) GetTrxByID(ctx *fiber.Ctx) error {
	c := ctx.Context()

	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	resUsecase, errUsecase := u.usecase.GetTrxByID(c, int64(id))

	if errUsecase != nil {
		return ctx.Status(errUsecase.Code).JSON(fiber.Map{
			"message": errUsecase.Message.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    resUsecase,
	})

}

// func (u *trxControllerImpl) UpdateTrx(ctx *fiber.Ctx) error {
// 	c := ctx.Context()

// 	id, err := strconv.Atoi(ctx.Params("id"))

// 	if err != nil {
// 		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": err.Error(),
// 		})
// 	}

// 	var trxData dto.TrxRequest

// 	if err := ctx.BodyParser(&trxData); err != nil {
// 		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": err.Error(),
// 		})
// 	}

// 	resUsecase, errUsecase := u.usecase.UpdateTrx(c, int64(id), trxData)

// 	if errUsecase != nil {
// 		return ctx.Status(errUsecase.Code).JSON(fiber.Map{
// 			"message": errUsecase.Message.Error(),
// 		})
// 	}

// 	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"message": "success",
// 		"data":    resUsecase,
// 	})

// }

func (u *trxControllerImpl) DeleteTrx(ctx *fiber.Ctx) error {
	c := ctx.Context()

	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errUsecase := u.usecase.DeleteTrx(c, int64(id))

	if errUsecase != nil {
		return ctx.Status(errUsecase.Code).JSON(fiber.Map{
			"message": errUsecase.Message.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})

}
