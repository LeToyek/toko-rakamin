package usecase

import (
	"context"
	"rakamin-final/internal/daos"
	"rakamin-final/internal/helper"
	"rakamin-final/internal/pkg/dto"
	"rakamin-final/internal/pkg/repository"

	"github.com/gofiber/fiber/v2"
)

type ProductPhotoUsecase interface {
	CreateProductPhoto(c context.Context, productPhoto dto.ProductPhotoRequest) (res dto.ProductPhotoResponse, err *helper.ErrorStruct)
	GetAllProductPhotos(c context.Context) (res []dto.ProductPhotoResponse, err *helper.ErrorStruct)
	GetProductPhotoByID(c context.Context, id int64) (res dto.ProductPhotoResponse, err *helper.ErrorStruct)
	UpdateProductPhoto(c context.Context, id int64, productPhoto dto.ProductPhotoRequest) (res dto.ProductPhotoResponse, err *helper.ErrorStruct)
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
		return res, &helper.ErrorStruct{
			Code:    fiber.StatusInternalServerError,
			Message: errRepo,
		}
	}

	for _, v := range resRepo {
		res = append(res, dto.ProductPhotoResponse{
			ID:  v.ID,
			Url: v.Url,
			Produk: dto.ProductResponse{
				ID:            v.Produk.ID,
				NamaProduk:    v.Produk.NamaProduk,
				Slug:          v.Produk.Slug,
				HargaReseller: v.Produk.HargaReseller,
				HargaKonsumen: v.Produk.HargaKonsumen,
				Stok:          v.Produk.Stok,
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
		Produk: dto.ProductResponse{
			ID:            resRepo.Produk.ID,
			NamaProduk:    resRepo.Produk.NamaProduk,
			Slug:          resRepo.Produk.Slug,
			HargaReseller: resRepo.Produk.HargaReseller,
			HargaKonsumen: resRepo.Produk.HargaKonsumen,
			Stok:          resRepo.Produk.Stok,
		},
	}

	return res, nil
}

func (u *productPhotoUsecaseImpl) UpdateProductPhoto(c context.Context, id int64, productPhoto dto.ProductPhotoRequest) (res dto.ProductPhotoResponse, err *helper.ErrorStruct) {

	if validateErr := helper.Validate.Struct(productPhoto); validateErr != nil {
		return res, &helper.ErrorStruct{
			Code:    fiber.StatusBadRequest,
			Message: validateErr,
		}
	}

	resRepo, errRepo := u.repo.UpdateProductPhoto(c, id, daos.FotoProduk{
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
