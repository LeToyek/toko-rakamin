package repository

import (
	"context"
	"rakamin-final/internal/daos"

	"gorm.io/gorm"
)

type ProductPhotoRepository interface {
	GetAllProductPhotos(ctx context.Context) (res []daos.FotoProduk, err error)
	GetProductPhotoByID(ctx context.Context, id int64) (res daos.FotoProduk, err error)
	CreateProductPhoto(ctx context.Context, productPhoto daos.FotoProduk) (res daos.FotoProduk, err error)
	UpdateProductPhoto(ctx context.Context, id int64, productPhoto daos.FotoProduk) (res daos.FotoProduk, err error)
	DeleteProductPhoto(ctx context.Context, id int64) (res daos.FotoProduk, err error)
}

type productPhotoRepositoryImpl struct {
	db *gorm.DB
}

func NewProductPhotoRepository(db *gorm.DB) *productPhotoRepositoryImpl {
	return &productPhotoRepositoryImpl{db}
}

func (r *productPhotoRepositoryImpl) GetAllProductPhotos(ctx context.Context) (res []daos.FotoProduk, err error) {
	db := r.db.Preload("Produk")

	if err := db.Where(&daos.FotoProduk{}).WithContext(ctx).Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (r *productPhotoRepositoryImpl) GetProductPhotoByID(ctx context.Context, id int64) (res daos.FotoProduk, err error) {

	db := r.db.Preload("Produk")

	if err := db.Where(&daos.FotoProduk{
		ID: id,
	}).WithContext(ctx).First(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (r *productPhotoRepositoryImpl) CreateProductPhoto(ctx context.Context, productPhoto daos.FotoProduk) (res daos.FotoProduk, err error) {

	db := r.db

	if err := db.WithContext(ctx).Create(&productPhoto).Error; err != nil {
		return res, err
	}

	return productPhoto, nil
}

func (r *productPhotoRepositoryImpl) UpdateProductPhoto(ctx context.Context, id int64, productPhoto daos.FotoProduk) (res daos.FotoProduk, err error) {

	db := r.db

	if err := db.Where(&daos.FotoProduk{
		ID: id,
	}).WithContext(ctx).Updates(&productPhoto).Error; err != nil {
		return res, err
	}

	return productPhoto, nil
}

func (r *productPhotoRepositoryImpl) DeleteProductPhoto(ctx context.Context, id int64) (res daos.FotoProduk, err error) {

	db := r.db

	if err := db.Where(&daos.FotoProduk{
		ID: id,
	}).WithContext(ctx).Delete(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}
