package controller

import (
	"rakamin-final/internal/pkg/dto"
	"rakamin-final/internal/pkg/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductController interface {
	CreateProduct(ctx *fiber.Ctx) error
	GetAllProducts(ctx *fiber.Ctx) error
	GetProductByID(ctx *fiber.Ctx) error
	UpdateProduct(ctx *fiber.Ctx) error
	DeleteProduct(ctx *fiber.Ctx) error
}

type productControllerImpl struct {
	usecase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) *productControllerImpl {
	return &productControllerImpl{
		usecase: usecase,
	}
}

func (u *productControllerImpl) CreateProduct(ctx *fiber.Ctx) error {
	c := ctx.Context()

	var product dto.ProductRequest

	if err := ctx.BodyParser(&product); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	res, err := u.usecase.CreateProduct(c, product)

	if err != nil {
		return ctx.Status(err.Code).JSON(fiber.Map{
			"message": err.Message.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
		"data":    res,
		"product": product,
	})

}

func (u *productControllerImpl) GetAllProducts(ctx *fiber.Ctx) error {
	c := ctx.Context()

	res, err := u.usecase.GetAllProducts(c, dto.ProductFilter{})

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

func (u *productControllerImpl) GetProductByID(ctx *fiber.Ctx) error {
	c := ctx.Context()

	id, _ := strconv.Atoi(ctx.Params("id"))

	res, err := u.usecase.GetProductByID(c, int64(id))

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

func (u *productControllerImpl) UpdateProduct(ctx *fiber.Ctx) error {
	c := ctx.Context()

	id, _ := strconv.Atoi(ctx.Params("id"))

	var product dto.ProductRequest

	if err := ctx.BodyParser(&product); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	res, err := u.usecase.UpdateProduct(c, int64(id), product)

	if err != nil {
		return ctx.Status(err.Code).JSON(fiber.Map{
			"message": err.Message.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    res,
		"product": product,
	})
}

func (u *productControllerImpl) DeleteProduct(ctx *fiber.Ctx) error {
	c := ctx.Context()

	id, _ := strconv.Atoi(ctx.Params("id"))

	err := u.usecase.DeleteProduct(c, int64(id))

	if err != nil {
		return ctx.Status(err.Code).JSON(fiber.Map{
			"message": err.Message.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})

}
