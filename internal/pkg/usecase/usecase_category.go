package usecase

import (
	"context"
	"rakamin-final/internal/daos"
	"rakamin-final/internal/helper"
	"rakamin-final/internal/pkg/dto"
	"rakamin-final/internal/pkg/repository"
)

type CategoryUsecase interface {
	GetAllCategories(ctx context.Context) (res []dto.Category, err *helper.ErrorStruct)
	GetCategoryByID(ctx context.Context, id int64) (res dto.CategoryResponse, err *helper.ErrorStruct)
	CreateCategory(ctx context.Context, category dto.CategoryRequest) (res dto.CategoryResponse, err *helper.ErrorStruct)
	UpdateCategory(ctx context.Context, id int64, category dto.CategoryRequest) (res dto.CategoryResponse, err *helper.ErrorStruct)
	DeleteCategory(ctx context.Context, id int64) (res dto.CategoryResponse, err *helper.ErrorStruct)
}

type categoryUsecaseImpl struct {
	repo repository.CategoryRepository
}

func NewCategoryUsecase(repo repository.CategoryRepository) *categoryUsecaseImpl {
	return &categoryUsecaseImpl{repo}
}

func (u *categoryUsecaseImpl) GetAllCategories(ctx context.Context) (res []dto.Category, err *helper.ErrorStruct) {
	resRepo, errRepo := u.repo.GetAllCategories(ctx)

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Code:    500,
			Message: errRepo,
		}
	}

	for _, v := range resRepo {
		res = append(res, dto.Category{
			ID:        v.ID,
			Nama:      v.NamaCategory,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}
	return res, nil
}

func (u *categoryUsecaseImpl) GetCategoryByID(ctx context.Context, id int64) (res dto.CategoryResponse, err *helper.ErrorStruct) {
	resRepo, errRepo := u.repo.GetCategoryByID(ctx, id)

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Code:    500,
			Message: errRepo,
		}
	}

	res = dto.CategoryResponse{
		ID:        resRepo.ID,
		Nama:      resRepo.NamaCategory,
		CreatedAt: resRepo.CreatedAt,
		UpdatedAt: resRepo.UpdatedAt,
	}
	return res, nil
}

func (u *categoryUsecaseImpl) CreateCategory(ctx context.Context, category dto.CategoryRequest) (res dto.CategoryResponse, err *helper.ErrorStruct) {
	resRepo, errRepo := u.repo.CreateCategory(ctx, daos.Category{
		NamaCategory: category.Nama,
	})

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Code:    500,
			Message: errRepo,
		}
	}

	res = dto.CategoryResponse{
		ID:        resRepo.ID,
		Nama:      resRepo.NamaCategory,
		CreatedAt: resRepo.CreatedAt,
		UpdatedAt: resRepo.UpdatedAt,
	}
	return res, nil
}

func (u *categoryUsecaseImpl) UpdateCategory(ctx context.Context, id int64, category dto.CategoryRequest) (res dto.CategoryResponse, err *helper.ErrorStruct) {
	resRepo, errRepo := u.repo.UpdateCategory(ctx, id, daos.Category{
		NamaCategory: category.Nama,
	})

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Code:    500,
			Message: errRepo,
		}
	}

	res = dto.CategoryResponse{
		ID:        resRepo.ID,
		Nama:      resRepo.NamaCategory,
		CreatedAt: resRepo.CreatedAt,
		UpdatedAt: resRepo.UpdatedAt,
	}
	return res, nil
}

func (u *categoryUsecaseImpl) DeleteCategory(ctx context.Context, id int64) (res dto.CategoryResponse, err *helper.ErrorStruct) {
	resRepo, errRepo := u.repo.DeleteCategory(ctx, id)

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Code:    500,
			Message: errRepo,
		}
	}

	res = dto.CategoryResponse{
		ID:        resRepo.ID,
		Nama:      resRepo.NamaCategory,
		CreatedAt: resRepo.CreatedAt,
		UpdatedAt: resRepo.UpdatedAt,
	}
	return res, nil
}
