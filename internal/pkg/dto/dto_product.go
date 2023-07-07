package dto

import "time"

type ProductFilter struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
}

type ProductResponse struct {
	ID            int64  `json:"id"`
	NamaProduk    string `json:"nama_produk" validate:"required,min=3,max=255"`
	Slug          string `json:"slug" validate:"required,min=3,max=255"`
	HargaReseller string `json:"harga_reseller"`
	HargaKonsumen string `json:"harga_konsumen"`
	Stok          int    `json:"stok"`
	Deskripsi     string `json:"deskripsi" validate:"required,min=3"`
	CreatedAt     time.Time
	UpdatedAt     time.Time

	// Toko        StoreResponse
	// Category    Category     `json:"foreignKey:CategoryID"`
	// FotoProduks []FotoProduk `json:"foreignKey:IdProduk"`
}

type ProductRequest struct {
	NamaProduk    string `json:"nama_produk" validate:"required,min=3,max=255"`
	Slug          string `json:"slug" validate:"required,min=3,max=255"`
	HargaReseller string `json:"harga_reseller"`
	HargaKonsumen string `json:"harga_konsumen"`
	Stok          int    `json:"stok"`
	Deskripsi     string `json:"deskripsi" validate:"required,min=3"`
	// IdToko        int64
	// Toko          Toko
	// CategoryID    int64
	// Category      Category     `json:"foreignKey:CategoryID"`
	// FotoProduks   []FotoProduk `json:"foreignKey:IdProduk"`
}
