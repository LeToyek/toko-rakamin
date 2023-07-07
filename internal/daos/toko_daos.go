package daos

import (
	"time"

	"gorm.io/gorm"
)

type (
	Toko struct {
		gorm.Model
		ID         int64 `gorm:"primaryKey;autoIncrement:true"`
		IdUser     int64
		User       User        `gorm:"foreignKey:IdUser"`
		NamaToko   string      `gorm:"size:255"`
		UrlFoto    string      `gorm:"size:255"`
		UpdatedAt  time.Time   `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP()"`
		CreatedAt  time.Time   `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP()"`
		Produks    []Produk    `gorm:"foreignKey:IdToko"`
		LogProduks []LogProduk `gorm:"foreignKey:IdToko"`
		DetailTrxs []DetailTrx `gorm:"foreignKey:IdToko"`
	}
	FilterToko struct {
		ID            int64
		Limit, Offset int
		NamaToko      string
	}
)
