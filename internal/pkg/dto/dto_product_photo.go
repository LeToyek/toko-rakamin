package dto

import "time"

type ProductPhotoRequest struct {
	IdProduk int64  `json:"id_produk" validate:"required"`
	Url      string `json:"url" validate:"required"`
}

type ProductPhotoResponse struct {
	ID        int64  `json:"id"`
	Url       string `json:"url"`
	Produk    ProductResponseForPhoto
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
type ProductResponseForProduct struct {
	ID  int64  `json:"id"`
	Url string `json:"url"`
}

type ProductPhotoRequestEdit struct {
	Url string `json:"url" validate:"required"`
}
