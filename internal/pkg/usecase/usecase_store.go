package usecase

import (
	"context"
	"errors"
	"rakamin-final/internal/daos"
	"rakamin-final/internal/helper"
	storeDTO "rakamin-final/internal/pkg/dto"
	"rakamin-final/internal/pkg/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type StoreUsecase interface {
	GetAllStores(ctx context.Context, params storeDTO.StoreFilter) (res []storeDTO.StoreResponse, err *helper.ErrorStruct)
	GetStoreByID(ctx context.Context, id int64) (res storeDTO.StoreResponse, err *helper.ErrorStruct)
	CreateStore(ctx context.Context, store storeDTO.StoreRequest) (res storeDTO.StoreResponse, err *helper.ErrorStruct)
	UpdateStore(ctx context.Context, id int64, Store storeDTO.StoreRequest) (res storeDTO.StoreResponse, err *helper.ErrorStruct)
	DeleteStore(ctx context.Context, id int64) *helper.ErrorStruct
}

type storeUsecaseImpl struct {
	repo repository.StoreRepository
}

func NewStoreUsecase(repo repository.StoreRepository) *storeUsecaseImpl {
	return &storeUsecaseImpl{
		repo: repo,
	}
}

func (u *storeUsecaseImpl) GetAllStores(ctx context.Context, params storeDTO.StoreFilter) (res []storeDTO.StoreResponse, err *helper.ErrorStruct) {
	if params.Limit < 1 {
		params.Limit = 10
	}
	if params.Page < 1 {
		params.Page = 0
	} else {
		params.Page = (params.Page - 1) * params.Limit
	}
	resRepo, errRepo := u.repo.GetAllStores(ctx, daos.FilterToko{
		ID:       int64(params.ID),
		Limit:    params.Limit,
		NamaToko: params.Name,
	})

	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code:    fiber.StatusNotFound,
			Message: errors.New("data not found"),
		}
	}
	for _, v := range resRepo {
		res = append(res, storeDTO.StoreResponse{
			ID:        v.ID,
			NamaToko:  v.NamaToko,
			UrlFoto:   v.UrlFoto,
			UpdatedAt: v.UpdatedAt,
			CreatedAt: v.CreatedAt,
		})
	}
	return res, nil
}

func (u *storeUsecaseImpl) GetStoreByID(ctx context.Context, id int64) (res storeDTO.StoreResponse, err *helper.ErrorStruct) {

	resRepo, errRepo := u.repo.GetStoreByID(ctx, id)

	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code:    fiber.StatusNotFound,
			Message: errors.New("data not found"),
		}
	}

	res = storeDTO.StoreResponse{
		ID:        resRepo.ID,
		NamaToko:  resRepo.NamaToko,
		UrlFoto:   resRepo.UrlFoto,
		UpdatedAt: resRepo.UpdatedAt,
		CreatedAt: resRepo.CreatedAt,
	}

	return res, nil

}

func (u *storeUsecaseImpl) CreateStore(ctx context.Context, store storeDTO.StoreRequest) (res storeDTO.StoreResponse, err *helper.ErrorStruct) {

	resRepo, errRepo := u.repo.CreateStore(ctx, daos.Toko{
		NamaToko: store.NamaToko,
		UrlFoto:  store.UrlFoto,
	})

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Code:    fiber.StatusConflict,
			Message: errRepo,
		}
	}

	res = storeDTO.StoreResponse{
		ID:        resRepo.ID,
		NamaToko:  resRepo.NamaToko,
		UrlFoto:   resRepo.UrlFoto,
		UpdatedAt: resRepo.UpdatedAt,
		CreatedAt: resRepo.CreatedAt,
	}

	return res, nil
}

func (u *storeUsecaseImpl) UpdateStore(ctx context.Context, id int64, store storeDTO.StoreRequest) (res storeDTO.StoreResponse, err *helper.ErrorStruct) {

	resRepo, errRepo := u.repo.UpdateStore(ctx, id, daos.Toko{
		NamaToko: store.NamaToko,
		UrlFoto:  store.UrlFoto,
	})

	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code:    fiber.StatusNotFound,
			Message: errors.New("data not found"),
		}
	}

	res = storeDTO.StoreResponse{
		ID:        resRepo.ID,
		NamaToko:  resRepo.NamaToko,
		UrlFoto:   resRepo.UrlFoto,
		UpdatedAt: resRepo.UpdatedAt,
		CreatedAt: resRepo.CreatedAt,
	}

	return res, nil
}

func (u *storeUsecaseImpl) DeleteStore(ctx context.Context, id int64) *helper.ErrorStruct {

	errRepo := u.repo.DeleteStore(ctx, id)

	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return &helper.ErrorStruct{
			Code:    fiber.StatusNotFound,
			Message: errors.New("data not found"),
		}
	}

	return nil
}
