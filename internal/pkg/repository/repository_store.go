package repository

import (
	"context"
	"fmt"
	"rakamin-final/internal/daos"
	"reflect"
	"strings"

	"gorm.io/gorm"
)

type StoreRepository interface {
	GetAllStores(ctx context.Context, params daos.FilterToko) (res []daos.Toko, err error)
	GetStoreByID(ctx context.Context, id int64) (res daos.Toko, err error)
	CreateStore(ctx context.Context, store daos.Toko) (res daos.Toko, err error)
	UpdateStore(ctx context.Context, id int64, store daos.Toko) (res daos.Toko, err error)
	DeleteStore(ctx context.Context, id int64) error
}

type storeRepositoryImpl struct {
	db *gorm.DB
}

func NewStoreRepository(db *gorm.DB) *storeRepositoryImpl {
	return &storeRepositoryImpl{db}
}

func (r *storeRepositoryImpl) GetAllStores(ctx context.Context, params daos.FilterToko) (res []daos.Toko, err error) {
	db := r.db

	structType := reflect.TypeOf(params)
	structValue := reflect.ValueOf(params)

	whereConditions := make([]string, 0)
	whereValues := make([]interface{}, 0)

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		value := structValue.Field(i)

		if value.Interface() != reflect.Zero(field.Type).Interface() {
			if field.Name != "Limit" && field.Name != "Offset" {
				whereConditions = append(whereConditions, fmt.Sprintf("%v like ?", field.Name))
				whereValues = append(whereValues, value.Interface())
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

func (r *storeRepositoryImpl) GetStoreByID(ctx context.Context, id int64) (res daos.Toko, err error) {
	if err := r.db.Where(&daos.Toko{
		ID: id,
	}).WithContext(ctx).First(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (r *storeRepositoryImpl) CreateStore(ctx context.Context, store daos.Toko) (res daos.Toko, err error) {
	db := r.db

	if err := db.Create(&store).WithContext(ctx).Error; err != nil {
		return res, err
	}

	return store, nil
}

func (r *storeRepositoryImpl) UpdateStore(ctx context.Context, id int64, store daos.Toko) (res daos.Toko, err error) {
	db := r.db

	if err := db.Where(&daos.Toko{
		ID: id,
	}).Updates(&store).WithContext(ctx).Error; err != nil {
		return res, err
	}

	return store, nil
}
func (r *storeRepositoryImpl) DeleteStore(ctx context.Context, id int64) error {

	db := r.db

	if err := db.Delete(&daos.Toko{}, id).WithContext(ctx).Error; err != nil {
		return err
	}

	return nil
}
