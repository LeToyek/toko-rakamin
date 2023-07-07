package controller

import (
	"rakamin-final/internal/pkg/dto"
	"rakamin-final/internal/pkg/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CategoryController interface {
	CreateCategory(ctx *fiber.Ctx) error
	GetAllCategories(ctx *fiber.Ctx) error
	GetCategoryByID(ctx *fiber.Ctx) error
	UpdateCategory(ctx *fiber.Ctx) error
	DeleteCategory(ctx *fiber.Ctx) error
}

type categoryControllerImpl struct {
	usecase usecase.CategoryUsecase
}

func NewCategoryController(usecase usecase.CategoryUsecase) *categoryControllerImpl {

	return &categoryControllerImpl{usecase}
}

func (u *categoryControllerImpl) CreateCategory(ctx *fiber.Ctx) error {
	c := ctx.Context()

	var category dto.CategoryRequest

	if err := ctx.BodyParser(&category); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	resUsecase, errUsecase := u.usecase.CreateCategory(c, category)

	if errUsecase != nil {
		return ctx.Status(errUsecase.Code).JSON(fiber.Map{
			"message": errUsecase.Message,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":  "success",
		"data":     resUsecase,
		"category": category,
	})

}

func (u *categoryControllerImpl) GetAllCategories(ctx *fiber.Ctx) error {
	c := ctx.Context()

	res, errUsecase := u.usecase.GetAllCategories(c)

	if errUsecase != nil {
		return ctx.Status(errUsecase.Code).JSON(fiber.Map{
			"message": errUsecase.Message,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    res,
	})

}

func (u *categoryControllerImpl) GetCategoryByID(ctx *fiber.Ctx) error {
	c := ctx.Context()

	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	res, errUsecase := u.usecase.GetCategoryByID(c, int64(id))

	if errUsecase != nil {
		return ctx.Status(errUsecase.Code).JSON(fiber.Map{
			"message": errUsecase.Message,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    res,
	})

}

func (u *categoryControllerImpl) UpdateCategory(ctx *fiber.Ctx) error {
	c := ctx.Context()

	var category dto.CategoryRequest

	if err := ctx.BodyParser(&category); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	id, _ := strconv.Atoi(ctx.Params("id"))

	res, errUsecase := u.usecase.UpdateCategory(c, int64(id), category)

	if errUsecase != nil {
		return ctx.Status(errUsecase.Code).JSON(fiber.Map{
			"message": errUsecase.Message,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":  "success",
		"data":     res,
		"category": category,
	})

}

func (u *categoryControllerImpl) DeleteCategory(ctx *fiber.Ctx) error {
	c := ctx.Context()

	id, _ := strconv.Atoi(ctx.Params("id"))

	res, errUsecase := u.usecase.DeleteCategory(c, int64(id))

	if errUsecase != nil {
		return ctx.Status(errUsecase.Code).JSON(fiber.Map{
			"message": errUsecase.Message,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    res,
	})

}
