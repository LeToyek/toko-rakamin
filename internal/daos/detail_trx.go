package daos

import (
	"time"
)

type DetailTrx struct {
	ID         int64 `gorm:"primaryKey;autoIncrement:true"`
	IdTrx      int64
	Trx        Trx `gorm:"foreignKey:IdTrx"`
	IdProduk   int64
	Produk     Produk `gorm:"foreignKey:IdProduk"`
	IdToko     int64
	Toko       Toko `gorm:"foreignKey:IdToko"`
	Kuantitas  int
	HargaTotal int
	UpdatedAt  time.Time `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP()"`
	CreatedAt  time.Time `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP()"`
}
