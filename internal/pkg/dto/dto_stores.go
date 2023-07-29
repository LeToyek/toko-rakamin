package dto

import "time"

type StoreFilter struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Limit  int    `json:"limit"`
	Page   int    `json:"page"`
	UserID int64  `json:"user_id"`
}

type StoreResponse struct {
	ID int64
	// User       User        `json:"foreignKey:IdUser"`
	NamaToko  string
	UrlFoto   string
	UpdatedAt time.Time
	CreatedAt time.Time
	User      UserResponse
	// Produks   []ProductResponseForStore `json:"foreignKey:IdToko"`
	// LogProduks []LogProduk `json:"foreignKey:IdToko"`
	// DetailTrxs []DetailTrx `json:"foreignKey:IdToko"`
}

type StoreRequest struct {
	NamaToko string `json:"nama_toko" validate:"required,min=3,max=255"`
	UrlFoto  string `json:"url_foto" validate:"required,min=3,max=255"`
}
