package controller

import (
	"fmt"
	"rakamin-final/internal/pkg/dto"
	"rakamin-final/internal/pkg/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductPhotoController interface {
	GetAllProductPhotos(c *fiber.Ctx) error
	GetProductPhotoByID(c *fiber.Ctx) error
	CreateProductPhoto(c *fiber.Ctx) error
	UpdateProductPhoto(c *fiber.Ctx) error
	DeleteProductPhoto(c *fiber.Ctx) error
}

type productPhotoControllerImpl struct {
	usecase usecase.ProductPhotoUsecase
}

func NewProductPhotoController(usecase usecase.ProductPhotoUsecase) *productPhotoControllerImpl {
	return &productPhotoControllerImpl{usecase}
}

func (p *productPhotoControllerImpl) GetAllProductPhotos(c *fiber.Ctx) error {
	res, err := p.usecase.GetAllProductPhotos(c.Context())
	if err != nil {
		return c.Status(err.Code).JSON(err)
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

func (p *productPhotoControllerImpl) GetProductPhotoByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	res, err := p.usecase.GetProductPhotoByID(c.Context(), int64(id))
	if err != nil {
		return c.Status(err.Code).JSON(err)
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

func (p *productPhotoControllerImpl) CreateProductPhoto(c *fiber.Ctx) error {
	var productPhotoRequest dto.ProductPhotoRequest
	if err := c.BodyParser(&productPhotoRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	res, err := p.usecase.CreateProductPhoto(c.Context(), productPhotoRequest)
	if err != nil {
		return c.Status(err.Code).JSON(err)
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": fmt.Sprintf("product photo with url %s successfully created", res.Url)})
}

func (p *productPhotoControllerImpl) UpdateProductPhoto(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var productPhotoRequest dto.ProductPhotoRequestEdit
	if err := c.BodyParser(&productPhotoRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	_, err := p.usecase.UpdateProductPhoto(c.Context(), int64(id), productPhotoRequest)
	if err != nil {
		return c.Status(err.Code).JSON(err)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})
}

func (p *productPhotoControllerImpl) DeleteProductPhoto(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	_, err := p.usecase.DeleteProductPhoto(c.Context(), int64(id))
	if err != nil {
		return c.Status(err.Code).JSON(err)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": fmt.Sprintf("product photo with id %d successfully deleted", id)})
}
