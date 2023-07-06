package daos

import (
	"time"
)

type (
	FotoProduk struct {
		ID        int64 `gorm:"primaryKey;autoIncrement:true"`
		IdProduk  int64
		Produk    Produk    `gorm:"foreignKey:IdProduk"`
		Url       string    `gorm:"size:255"`
		UpdatedAt time.Time `gorm:"type:date;default:CURRENT_TIMESTAMP()"`
		CreatedAt time.Time `gorm:"type:date;default:CURRENT_TIMESTAMP()"`
	}
)
