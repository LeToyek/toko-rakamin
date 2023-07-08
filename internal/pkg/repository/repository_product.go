package repository

import (
	"context"
	"fmt"
	"rakamin-final/internal/daos"
	"reflect"
	"strings"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAllProducts(ctx context.Context, params daos.FilterProduk) (res []daos.Produk, err error)
	GetProductByID(ctx context.Context, id int64) (res daos.Produk, err error)
	CreateProduct(ctx context.Context, product daos.Produk) (res daos.Produk, err error)
	UpdateProduct(ctx context.Context, id int64, product daos.Produk) (res daos.Produk, err error)
	DeleteProduct(ctx context.Context, id int64) error
}

type productRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepositoryImpl {
	return &productRepositoryImpl{db}
}

func (r *productRepositoryImpl) GetAllProducts(ctx context.Context, params daos.FilterProduk) (res []daos.Produk, err error) {
	db := r.db.Preload("Toko").Preload("Category").Preload("FotoProduks").Limit(params.Limit).Offset(params.Offset)
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
	fmt.Println(res)
	return res, nil
}

func (r *productRepositoryImpl) GetProductByID(ctx context.Context, id int64) (res daos.Produk, err error) {
	err = r.db.Preload("Toko").Preload("Category").Preload("FotoProduks").Where("id = ?", id).WithContext(ctx).Find(&res).Error

	if err != nil {
		return res, err
	}

	fmt.Println(res.Toko.NamaToko)

	return res, nil
}

func (r *productRepositoryImpl) CreateProduct(ctx context.Context, product daos.Produk) (res daos.Produk, err error) {
	if err := r.db.Create(&product).WithContext(ctx).Error; err != nil {
		return res, err
	}
	return product, nil
}

func (r *productRepositoryImpl) UpdateProduct(ctx context.Context, id int64, product daos.Produk) (res daos.Produk, err error) {
	if err := r.db.Model(&daos.Produk{}).Where("id = ?", id).Updates(product).WithContext(ctx).Error; err != nil {
		return res, err
	}
	return product, nil
}

func (r *productRepositoryImpl) DeleteProduct(ctx context.Context, id int64) error {
	if err := r.db.Where("id = ?", id).Delete(&daos.Produk{}).WithContext(ctx).Error; err != nil {
		return err
	}
	return nil
}
