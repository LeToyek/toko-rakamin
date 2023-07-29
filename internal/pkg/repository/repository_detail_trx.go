package repository

import (
	"context"
	"rakamin-final/internal/daos"

	"gorm.io/gorm"
)

type DetailTrxRepository interface {
	CreateDetailTrx(ctx context.Context, detailTrx daos.DetailTrx) (res daos.DetailTrx, err error)
	GetDetailTrxByID(ctx context.Context, id int64) (res daos.DetailTrx, err error)
}

type detailTrxRepositoryImpl struct {
	db *gorm.DB
}

func NewDetailTrxRepository(db *gorm.DB) *detailTrxRepositoryImpl {
	return &detailTrxRepositoryImpl{db}
}

func (r *detailTrxRepositoryImpl) CreateDetailTrx(ctx context.Context, detailTrx daos.DetailTrx) (res daos.DetailTrx, err error) {

	if err := r.db.WithContext(ctx).Create(&detailTrx).Error; err != nil {
		return res, err
	}

	return detailTrx, nil
}
func (r *detailTrxRepositoryImpl) GetDetailTrxByID(ctx context.Context, id int64) (res daos.DetailTrx, err error) {
	if err := r.db.Where(&daos.DetailTrx{
		ID: id,
	}).WithContext(ctx).First(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}
