package daos

import (
	"time"
)

type (
	LogProduk struct {
		ID            int64 `gorm:"primaryKey;autoIncrement:true"`
		IdProduk      int64
		Produk        Produk    `gorm:"foreignKey:IdProduk"`
		NamaProduk    string    `gorm:"size:255"`
		Slug          string    `gorm:"size:255"`
		HargaReseller string    `gorm:"size:255"`
		HargaKonsumen string    `gorm:"size:255"`
		Deskripsi     string    `gorm:"type:text"`
		CreatedAt     time.Time `gorm:"type:date;default:CURRENT_TIMESTAMP()"`
		UpdatedAt     time.Time `gorm:"type:date;default:CURRENT_TIMESTAMP()"`
		IdToko        int64
		Toko          Toko `gorm:"foreignKey:IdToko"`
		CategoryID    int64
		Category      Category `gorm:"foreignKey:CategoryID"`
	}
)
