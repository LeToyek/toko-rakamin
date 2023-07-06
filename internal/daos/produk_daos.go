package daos

import (
	"time"
)

type (
	Produk struct {
		ID            int64  `gorm:"primaryKey;autoIncrement:true"`
		NamaProduk    string `gorm:"size:255"`
		Slug          string `gorm:"size:255"`
		HargaReseller string `gorm:"size:255"`
		HargaKonsumen string `gorm:"size:255"`
		Stok          int
		Deskripsi     string    `gorm:"type:text"`
		CreatedAt     time.Time `gorm:"type:date;default:CURRENT_TIMESTAMP()"`
		UpdatedAt     time.Time `gorm:"type:date;default:CURRENT_TIMESTAMP()"`
		IdToko        int64
		Toko          Toko `gorm:"foreignKey:IdToko"`
		CategoryID    int64
		Category      Category     `gorm:"foreignKey:CategoryID"`
		FotoProduks   []FotoProduk `gorm:"foreignKey:IdProduk"`
	}
)
