package repository

import (
	"context"
	"fmt"
	"rakamin-final/internal/daos"
	"reflect"
	"strings"

	"gorm.io/gorm"
)

type AddressRepository interface {
	GetAllAddress(ctx context.Context, params daos.FilterAlamat) (res []daos.Alamat, err error)
	GetAddressByID(ctx context.Context, id int64) (res daos.Alamat, err error)
	CreateAddress(ctx context.Context, address daos.Alamat) (res daos.Alamat, err error)
	UpdateAddress(ctx context.Context, id int64, address daos.Alamat) (res daos.Alamat, err error)
	DeleteAddress(ctx context.Context, id int64) error
}

type addressRepositoryImpl struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) *addressRepositoryImpl {
	return &addressRepositoryImpl{db}
}

func (r *addressRepositoryImpl) GetAllAddress(ctx context.Context, params daos.FilterAlamat) (res []daos.Alamat, err error) {
	db := r.db

	structType := reflect.TypeOf(params)
	structValue := reflect.ValueOf(params)

	whereConditions := make([]string, 0)
	whereValues := make([]interface{}, 0)

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		value := structValue.Field(i)

		if value.Interface() != reflect.Zero(field.Type).Interface() {
			whereConditions = append(whereConditions, fmt.Sprintf("%v like ?", field.Name))
			whereValues = append(whereValues, value.Interface())
		}
	}
	if len(whereConditions) > 0 {
		query := strings.Join(whereConditions, " OR ")
		err := db.Where(query, whereValues...).WithContext(ctx).Find(&res).Error
		if err != nil {
			return res, err
		}
	}

	return res, nil
}

func (r *addressRepositoryImpl) GetAddressByID(ctx context.Context, id int64) (res daos.Alamat, err error) {
	if err := r.db.Where(&daos.Alamat{
		ID: id,
	}).WithContext(ctx).First(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (r *addressRepositoryImpl) CreateAddress(ctx context.Context, address daos.Alamat) (res daos.Alamat, err error) {

	if err := r.db.WithContext(ctx).Create(&address).Error; err != nil {
		return res, err
	}
	return address, nil
}

func (r *addressRepositoryImpl) UpdateAddress(ctx context.Context, id int64, address daos.Alamat) (res daos.Alamat, err error) {

	if err := r.db.WithContext(ctx).Model(&daos.Alamat{}).Where("id = ?", id).Updates(&address).Error; err != nil {
		return res, err
	}

	return address, nil
}

func (r *addressRepositoryImpl) DeleteAddress(ctx context.Context, id int64) error {
	if err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&daos.Alamat{}).Error; err != nil {
		return err
	}

	return nil
}
