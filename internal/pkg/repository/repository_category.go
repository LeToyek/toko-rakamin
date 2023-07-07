package repository

import (
	"context"
	"rakamin-final/internal/daos"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetAllCategories(ctx context.Context) (res []daos.Category, err error)
	GetCategoryByID(ctx context.Context, id int64) (res daos.Category, err error)
	CreateCategory(ctx context.Context, category daos.Category) (res daos.Category, err error)
	UpdateCategory(ctx context.Context, id int64, category daos.Category) (res daos.Category, err error)
	DeleteCategory(ctx context.Context, id int64) (res daos.Category, err error)
}

type categoryRepositoryImpl struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepositoryImpl {
	return &categoryRepositoryImpl{db}
}

func (r *categoryRepositoryImpl) GetAllCategories(ctx context.Context) (res []daos.Category, err error) {
	db := r.db

	if err := db.Where(&daos.Category{}).WithContext(ctx).Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (r *categoryRepositoryImpl) GetCategoryByID(ctx context.Context, id int64) (res daos.Category, err error) {
	if err := r.db.Where(&daos.Category{
		ID: id,
	}).WithContext(ctx).First(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (r *categoryRepositoryImpl) CreateCategory(ctx context.Context, category daos.Category) (res daos.Category, err error) {

	if err := r.db.WithContext(ctx).Create(&category).Error; err != nil {
		return res, err
	}

	return category, nil
}

func (r *categoryRepositoryImpl) UpdateCategory(ctx context.Context, id int64, category daos.Category) (res daos.Category, err error) {

	if err := r.db.Where(&daos.Category{
		ID: id,
	}).WithContext(ctx).Updates(&category).Error; err != nil {
		return res, err
	}

	return category, nil
}

func (r *categoryRepositoryImpl) DeleteCategory(ctx context.Context, id int64) (res daos.Category, err error) {
	if err := r.db.Where(&daos.Category{
		ID: id,
	}).WithContext(ctx).Delete(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}
