package usecase

import (
	"context"
	"errors"
	"rakamin-final/internal/daos"
	"rakamin-final/internal/helper"
	"rakamin-final/internal/pkg/dto"
	"rakamin-final/internal/pkg/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ProductPhotoUsecase interface {
	CreateProductPhoto(c context.Context, productPhoto dto.ProductPhotoRequest) (res dto.ProductPhotoResponse, err *helper.ErrorStruct)
	GetAllProductPhotos(c context.Context) (res []dto.ProductPhotoResponse, err *helper.ErrorStruct)
	GetProductPhotoByID(c context.Context, id int64) (res dto.ProductPhotoResponse, err *helper.ErrorStruct)
	UpdateProductPhoto(c context.Context, id int64, productPhoto dto.ProductPhotoRequestEdit) (res dto.ProductPhotoResponse, err *helper.ErrorStruct)
	DeleteProductPhoto(c context.Context, id int64) (res dto.ProductPhotoResponse, err *helper.ErrorStruct)
}

type productPhotoUsecaseImpl struct {
	repo repository.ProductPhotoRepository
}

func NewProductPhotoUsecase(repo repository.ProductPhotoRepository) *productPhotoUsecaseImpl {
	return &productPhotoUsecaseImpl{repo}
}

func (u *productPhotoUsecaseImpl) CreateProductPhoto(c context.Context, productPhoto dto.ProductPhotoRequest) (res dto.ProductPhotoResponse, err *helper.ErrorStruct) {

	if validateErr := helper.Validate.Struct(productPhoto); validateErr != nil {
		return res, &helper.ErrorStruct{
			Code:    fiber.StatusBadRequest,
			Message: validateErr,
		}
	}

	resRepo, errRepo := u.repo.CreateProductPhoto(c, daos.FotoProduk{
		IdProduk: productPhoto.IdProduk,
		Url:      productPhoto.Url,
	})

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Code:    fiber.StatusInternalServerError,
			Message: errRepo,
		}
	}

	res = dto.ProductPhotoResponse{
		ID:        resRepo.ID,
		Url:       resRepo.Url,
		UpdatedAt: resRepo.UpdatedAt,
		CreatedAt: resRepo.CreatedAt,
	}

	return res, nil
}

func (u *productPhotoUsecaseImpl) GetAllProductPhotos(c context.Context) (res []dto.ProductPhotoResponse, err *helper.ErrorStruct) {

	resRepo, errRepo := u.repo.GetAllProductPhotos(c)

	if errRepo != nil {
		if errors.Is(errRepo, gorm.ErrRecordNotFound) {
			return res, &helper.ErrorStruct{
				Code:    fiber.StatusNotFound,
				Message: errors.New("data not found"),
			}
		}
	}

	for _, v := range resRepo {
		res = append(res, dto.ProductPhotoResponse{
			ID:  v.ID,
			Url: v.Url,
			Produk: dto.ProductResponseForPhoto{
				ID:            v.Produk.ID,
				NamaProduk:    v.Produk.NamaProduk,
				HargaReseller: v.Produk.HargaReseller,
				HargaKonsumen: v.Produk.HargaKonsumen,
				Stok:          v.Produk.Stok,
				Deskripsi:     v.Produk.Deskripsi,
			},
		})
	}

	return res, nil
}

func (u *productPhotoUsecaseImpl) GetProductPhotoByID(c context.Context, id int64) (res dto.ProductPhotoResponse, err *helper.ErrorStruct) {

	resRepo, errRepo := u.repo.GetProductPhotoByID(c, id)

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Code:    fiber.StatusInternalServerError,
			Message: errRepo,
		}
	}

	res = dto.ProductPhotoResponse{
		ID:  resRepo.ID,
		Url: resRepo.Url,
		Produk: dto.ProductResponseForPhoto{
			ID:            resRepo.Produk.ID,
			NamaProduk:    resRepo.Produk.NamaProduk,
			HargaReseller: resRepo.Produk.HargaReseller,
			HargaKonsumen: resRepo.Produk.HargaKonsumen,
			Stok:          resRepo.Produk.Stok,
			Deskripsi:     resRepo.Produk.Deskripsi,
		},
	}

	return res, nil
}

func (u *productPhotoUsecaseImpl) UpdateProductPhoto(c context.Context, id int64, productPhoto dto.ProductPhotoRequestEdit) (res dto.ProductPhotoResponse, err *helper.ErrorStruct) {

	if validateErr := helper.Validate.Struct(productPhoto); validateErr != nil {
		return res, &helper.ErrorStruct{
			Code:    fiber.StatusBadRequest,
			Message: validateErr,
		}
	}

	resRepo, errRepo := u.repo.UpdateProductPhoto(c, id, daos.FotoProduk{
		Url: productPhoto.Url,
	})

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Code:    fiber.StatusInternalServerError,
			Message: errRepo,
		}
	}

	res = dto.ProductPhotoResponse{
		ID:        resRepo.ID,
		Url:       resRepo.Url,
		UpdatedAt: resRepo.UpdatedAt,
		CreatedAt: resRepo.CreatedAt,
	}

	return res, nil
}

func (u *productPhotoUsecaseImpl) DeleteProductPhoto(c context.Context, id int64) (res dto.ProductPhotoResponse, err *helper.ErrorStruct) {

	resRepo, errRepo := u.repo.DeleteProductPhoto(c, id)

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Code:    fiber.StatusInternalServerError,
			Message: errRepo,
		}
	}

	res = dto.ProductPhotoResponse{
		ID:        resRepo.ID,
		Url:       resRepo.Url,
		UpdatedAt: resRepo.UpdatedAt,
		CreatedAt: resRepo.CreatedAt,
	}

	return res, nil
}
