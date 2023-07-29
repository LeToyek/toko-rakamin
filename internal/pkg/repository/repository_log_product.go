package repository

import (
	"context"
	"rakamin-final/internal/daos"

	"gorm.io/gorm"
)

type LogProductRepository interface {
	CreateLogProduct(ctx context.Context, logProduct daos.LogProduk) (res daos.LogProduk, err error)
	GetAllLogProducts(ctx context.Context) (res []daos.LogProduk, err error)
	GetLogProductByID(ctx context.Context, id int64) (res daos.LogProduk, err error)
}
type logProductRepositoryImpl struct {
	db *gorm.DB
}

func NewLogProductRepository(db *gorm.DB) *logProductRepositoryImpl {
	return &logProductRepositoryImpl{db}
}
func (r *logProductRepositoryImpl) CreateLogProduct(ctx context.Context, logProduct daos.LogProduk) (res daos.LogProduk, err error) {
	if err := r.db.WithContext(ctx).Create(&logProduct).Error; err != nil {
		return res, err
	}
	return logProduct, nil
}
func (r *logProductRepositoryImpl) GetAllLogProducts(ctx context.Context) (res []daos.LogProduk, err error) {
	db := r.db
	if err := db.Where(&daos.LogProduk{}).WithContext(ctx).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}
func (r *logProductRepositoryImpl) GetLogProductByID(ctx context.Context, id int64) (res daos.LogProduk, err error) {
	if err := r.db.Where(&daos.LogProduk{
		ID: id,
	}).WithContext(ctx).First(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}
