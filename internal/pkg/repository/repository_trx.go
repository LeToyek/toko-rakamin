package repository

import (
	"context"
	"fmt"
	"rakamin-final/internal/daos"
	"rakamin-final/internal/utils"
	"reflect"
	"strings"

	"gorm.io/gorm"
)

type TrxRepository interface {
	GetAllTrxes(ctx context.Context, params daos.FilterTrx) (res []daos.Trx, err error)
	GetTrxByID(ctx context.Context, id int64) (res daos.Trx, err error)
	CreateTrx(ctx context.Context, trx daos.Trx, detailTrxes []daos.DetailTrx) (res daos.Trx, err error)
	UpdateTrx(ctx context.Context, id int64, trx daos.Trx) (res daos.Trx, err error)
	DeleteTrx(ctx context.Context, id int64) error
}

type trxRepositoryImpl struct {
	db *gorm.DB
}

func NewTrxRepository(db *gorm.DB) *trxRepositoryImpl {
	return &trxRepositoryImpl{db}
}

func (r *trxRepositoryImpl) GetAllTrxes(ctx context.Context, params daos.FilterTrx) (res []daos.Trx, err error) {
	db := r.db.Preload("DetailTrxes").Preload("DetailTrxes.Produk").Preload("Alamat").Preload("Alamat.User").Preload("Toko").Preload("Toko.User").Limit(params.Limit).Offset(params.Offset)

	structType := reflect.TypeOf(params)
	structValue := reflect.ValueOf(params)

	whereConditions := make([]string, 0)
	whereValues := make([]interface{}, 0)

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		value := structValue.Field(i)

		if value.Interface() != reflect.Zero(field.Type).Interface() {
			if field.Name != "Limit" && field.Name != "Offset" {
				camelCaseName := utils.GenerateSlugCamelCase(field.Name)
				whereConditions = append(whereConditions, fmt.Sprintf("%v like ?", camelCaseName))
				whereValues = append(whereValues, fmt.Sprintf("%%%v%%", value.Interface()))
			}
		}
	}
	if len(whereConditions) > 0 {
		query := strings.Join(whereConditions, " OR ")
		err := db.Where(query, whereValues...).WithContext(ctx).Find(&res).Error
		if err != nil {
			return res, err
		}
	} else {
		err := db.Find(&res).Error
		if err != nil {
			return res, err
		}
	}
	return res, nil
}

func (r *trxRepositoryImpl) GetTrxByID(ctx context.Context, id int64) (res daos.Trx, err error) {
	db := r.db.Preload("DetailTrxes").Preload("DetailTrxes.Produk").Preload("Alamat").Preload("Alamat.User").Preload("Toko").Preload("Toko.User")
	if err := db.Where(&daos.Trx{
		ID: id,
	}).WithContext(ctx).First(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (r *trxRepositoryImpl) CreateTrx(ctx context.Context, trx daos.Trx, detailTrxes []daos.DetailTrx) (res daos.Trx, err error) {
	db := r.db.Preload("DetailTrxes").Preload("DetailTrxes.Produk").Preload("Alamat").Preload("Alamat.User").Preload("Toko").Preload("Toko.User")
	err = db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&trx).WithContext(ctx).Error; err != nil {
			return err
		}
		for _, detailTrx := range detailTrxes {
			detailTrx.IdTrx = trx.ID
			if err := tx.Create(&detailTrx).WithContext(ctx).Error; err != nil {
				return err
			}
		}
		return err
	})

	res = trx

	return res, nil
}

func (r *trxRepositoryImpl) UpdateTrx(ctx context.Context, id int64, trx daos.Trx) (res daos.Trx, err error) {
	if err := r.db.Model(&daos.Trx{
		ID: id,
	}).Updates(trx).WithContext(ctx).Error; err != nil {
		return res, err
	}

	return trx, nil
}

func (r *trxRepositoryImpl) DeleteTrx(ctx context.Context, id int64) error {
	if err := r.db.Delete(&daos.Trx{
		ID: id,
	}).WithContext(ctx).Error; err != nil {
		return err
	}

	return nil
}
