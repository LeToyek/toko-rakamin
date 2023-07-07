package usecase

import (
	"context"
	"rakamin-final/internal/daos"
	"rakamin-final/internal/helper"
	ProductDTO "rakamin-final/internal/pkg/dto"
	"rakamin-final/internal/pkg/repository"
	"rakamin-final/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type ProductUsecase interface {
	GetAllProducts(ctx context.Context, params ProductDTO.ProductFilter) (res []ProductDTO.ProductResponse, err *helper.ErrorStruct)
	GetProductByID(ctx context.Context, id int64) (res ProductDTO.ProductResponse, err *helper.ErrorStruct)
	CreateProduct(ctx context.Context, params ProductDTO.ProductRequest) (res ProductDTO.ProductResponse, err *helper.ErrorStruct)
	UpdateProduct(ctx context.Context, id int64, params ProductDTO.ProductRequestUpdate) (res ProductDTO.ProductResponse, err *helper.ErrorStruct)
	DeleteProduct(ctx context.Context, id int64) *helper.ErrorStruct
}

type productUsecaseImpl struct {
	repo repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) *productUsecaseImpl {
	return &productUsecaseImpl{
		repo: repo,
	}
}

func (u *productUsecaseImpl) GetAllProducts(ctx context.Context, params ProductDTO.ProductFilter) (res []ProductDTO.ProductResponse, err *helper.ErrorStruct) {
	if params.Limit < 1 {
		params.Limit = 10
	}
	if params.Page < 1 {
		params.Page = 0
	} else {
		params.Page = (params.Page - 1) * params.Limit
	}
	resRepo, errRepo := u.repo.GetAllProducts(ctx, daos.FilterProduk{
		ID:         int64(params.ID),
		Limit:      params.Limit,
		NamaProduk: params.Name,
	})

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Code:    500,
			Message: errRepo,
		}
	}
	for _, v := range resRepo {
		res = append(res, ProductDTO.ProductResponse{
			ID:            v.ID,
			NamaProduk:    v.NamaProduk,
			Slug:          v.Slug,
			HargaReseller: v.HargaReseller,
			HargaKonsumen: v.HargaKonsumen,
			Stok:          v.Stok,
			Deskripsi:     v.Deskripsi,
		})
	}

	return res, nil

}

func (u *productUsecaseImpl) GetProductByID(ctx context.Context, id int64) (res ProductDTO.ProductResponse, err *helper.ErrorStruct) {
	resRepo, errRepo := u.repo.GetProductByID(ctx, id)
	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Code:    500,
			Message: errRepo,
		}
	}
	res = ProductDTO.ProductResponse{
		ID:            resRepo.ID,
		NamaProduk:    resRepo.NamaProduk,
		Slug:          resRepo.Slug,
		HargaReseller: resRepo.HargaReseller,
		HargaKonsumen: resRepo.HargaKonsumen,
		Stok:          resRepo.Stok,
		Deskripsi:     resRepo.Deskripsi,
	}

	return res, nil
}

func (u *productUsecaseImpl) CreateProduct(ctx context.Context, params ProductDTO.ProductRequest) (res ProductDTO.ProductResponse, err *helper.ErrorStruct) {
	if err := helper.Validate.Struct(params); err != nil {
		return res, &helper.ErrorStruct{
			Code:    fiber.StatusBadRequest,
			Message: err,
		}
	}
	resRepo, errRepo := u.repo.CreateProduct(ctx, daos.Produk{
		NamaProduk:    params.NamaProduk,
		Slug:          utils.GenerateSlug(params.NamaProduk),
		HargaReseller: params.HargaReseller,
		HargaKonsumen: params.HargaKonsumen,
		Stok:          params.Stok,
		Deskripsi:     params.Deskripsi,
		IdToko:        params.TokoID,
		CategoryID:    params.CategoryID,
	})
	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Code:    500,
			Message: errRepo,
		}
	}
	res = ProductDTO.ProductResponse{
		ID:            resRepo.ID,
		NamaProduk:    resRepo.NamaProduk,
		Slug:          resRepo.Slug,
		HargaReseller: resRepo.HargaReseller,
		HargaKonsumen: resRepo.HargaKonsumen,
		Stok:          resRepo.Stok,
		Deskripsi:     resRepo.Deskripsi,
	}

	return res, nil
}

func (u *productUsecaseImpl) UpdateProduct(ctx context.Context, id int64, params ProductDTO.ProductRequestUpdate) (res ProductDTO.ProductResponse, err *helper.ErrorStruct) {
	if err := helper.Validate.Struct(params); err != nil {
		return res, &helper.ErrorStruct{
			Code:    fiber.StatusBadRequest,
			Message: err,
		}
	}
	resRepo, errRepo := u.repo.UpdateProduct(ctx, id, daos.Produk{
		NamaProduk:    params.NamaProduk,
		Slug:          utils.GenerateSlug(params.NamaProduk),
		HargaReseller: params.HargaReseller,
		HargaKonsumen: params.HargaKonsumen,
		Stok:          params.Stok,
		Deskripsi:     params.Deskripsi,
		IdToko:        params.TokoID,
		CategoryID:    params.CategoryID,
	})
	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Code:    500,
			Message: errRepo,
		}
	}
	res = ProductDTO.ProductResponse{
		ID:            resRepo.ID,
		NamaProduk:    resRepo.NamaProduk,
		Slug:          resRepo.Slug,
		HargaReseller: resRepo.HargaReseller,
		HargaKonsumen: resRepo.HargaKonsumen,
		Stok:          resRepo.Stok,
		Deskripsi:     resRepo.Deskripsi,
	}

	return res, nil
}

func (u *productUsecaseImpl) DeleteProduct(ctx context.Context, id int64) *helper.ErrorStruct {
	errRepo := u.repo.DeleteProduct(ctx, id)
	if errRepo != nil {
		return &helper.ErrorStruct{
			Code:    500,
			Message: errRepo,
		}
	}
	return nil
}
