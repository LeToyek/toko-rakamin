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
	repoProduct    repository.ProductRepository
	repoLogProduct repository.LogProductRepository
}

func NewProductUsecase(repoProduct repository.ProductRepository, repoLogProduct repository.LogProductRepository) *productUsecaseImpl {
	return &productUsecaseImpl{
		repoProduct:    repoProduct,
		repoLogProduct: repoLogProduct,
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
	resRepo, errRepo := u.repoProduct.GetAllProducts(ctx, daos.FilterProduk{
		ID:         int64(params.ID),
		Limit:      params.Limit,
		Offset:     params.Page,
		NamaProduk: params.Name,
	})

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Code:    500,
			Message: errRepo,
		}
	}
	for i, v := range resRepo {
		var productPhotos []ProductDTO.ProductResponseForProduct
		for _, x := range resRepo[i].FotoProduks {
			productPhotos = append(productPhotos, ProductDTO.ProductResponseForProduct{
				ID:  x.ID,
				Url: x.Url})
		}
		res = append(res, ProductDTO.ProductResponse{
			ID:            v.ID,
			NamaProduk:    v.NamaProduk,
			Slug:          v.Slug,
			HargaReseller: v.HargaReseller,
			HargaKonsumen: v.HargaKonsumen,
			Stok:          v.Stok,
			Deskripsi:     v.Deskripsi,
			Toko: ProductDTO.StoreResponse{
				NamaToko:  v.Toko.NamaToko,
				UrlFoto:   v.Toko.UrlFoto,
				ID:        v.Toko.ID,
				CreatedAt: v.Toko.CreatedAt,
				UpdatedAt: v.Toko.UpdatedAt,
			},
			Category: ProductDTO.CategoryResponse{
				ID:        v.Category.ID,
				Nama:      v.Category.NamaCategory,
				CreatedAt: v.Category.CreatedAt,
				UpdatedAt: v.Category.UpdatedAt,
			},
			FotoProduks: productPhotos})

	}

	return res, nil

}

func (u *productUsecaseImpl) GetProductByID(ctx context.Context, id int64) (res ProductDTO.ProductResponse, err *helper.ErrorStruct) {
	resRepo, errRepo := u.repoProduct.GetProductByID(ctx, id)
	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Code:    500,
			Message: errRepo,
		}
	}
	var productPhotos []ProductDTO.ProductResponseForProduct
	for _, v := range resRepo.FotoProduks {
		productPhotos = append(productPhotos, ProductDTO.ProductResponseForProduct{
			ID:  v.ID,
			Url: v.Url,
		})
	}
	res = ProductDTO.ProductResponse{
		ID:            resRepo.ID,
		NamaProduk:    resRepo.NamaProduk,
		Slug:          resRepo.Slug,
		HargaReseller: resRepo.HargaReseller,
		HargaKonsumen: resRepo.HargaKonsumen,
		Stok:          resRepo.Stok,
		Deskripsi:     resRepo.Deskripsi,
		Toko: ProductDTO.StoreResponse{
			NamaToko:  resRepo.Toko.NamaToko,
			UrlFoto:   resRepo.Toko.UrlFoto,
			ID:        resRepo.Toko.ID,
			CreatedAt: resRepo.Toko.CreatedAt,
			UpdatedAt: resRepo.Toko.UpdatedAt,
		},
		Category: ProductDTO.CategoryResponse{
			ID:        resRepo.Category.ID,
			Nama:      resRepo.Category.NamaCategory,
			CreatedAt: resRepo.Category.CreatedAt,
			UpdatedAt: resRepo.Category.UpdatedAt,
		},
		FotoProduks: productPhotos,
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

	resRepo, errRepo := u.repoProduct.CreateProduct(ctx, daos.Produk{
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
	_, resError := u.repoLogProduct.CreateLogProduct(ctx, daos.LogProduk{
		IdProduk:      resRepo.ID,
		NamaProduk:    resRepo.NamaProduk,
		Slug:          resRepo.Slug,
		HargaReseller: resRepo.HargaReseller,
		HargaKonsumen: resRepo.HargaKonsumen,
		Deskripsi:     resRepo.Deskripsi,
		IdToko:        resRepo.IdToko,
		CategoryID:    resRepo.CategoryID,
	})
	if resError != nil {
		return res, &helper.ErrorStruct{
			Code:    500,
			Message: resError,
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
	resRepo, errRepo := u.repoProduct.UpdateProduct(ctx, id, daos.Produk{
		NamaProduk:    params.NamaProduk,
		Slug:          utils.GenerateSlug(params.NamaProduk),
		HargaReseller: params.HargaReseller,
		HargaKonsumen: params.HargaKonsumen,
		Stok:          params.Stok,
		Deskripsi:     params.Deskripsi,
		IdToko:        params.TokoID,
		CategoryID:    params.CategoryID,
	})
	_, resError := u.repoLogProduct.CreateLogProduct(ctx, daos.LogProduk{
		IdProduk:      resRepo.ID,
		NamaProduk:    resRepo.NamaProduk,
		Slug:          resRepo.Slug,
		HargaReseller: resRepo.HargaReseller,
		HargaKonsumen: resRepo.HargaKonsumen,
		Deskripsi:     resRepo.Deskripsi,
		IdToko:        resRepo.IdToko,
		CategoryID:    resRepo.CategoryID,
	})
	if resError != nil {
		return res, &helper.ErrorStruct{
			Code:    500,
			Message: resError,
		}
	}

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
	errRepo := u.repoProduct.DeleteProduct(ctx, id)
	if errRepo != nil {
		return &helper.ErrorStruct{
			Code:    500,
			Message: errRepo,
		}
	}
	return nil
}
