package daos

import (
	"time"
)

type DetailTrx struct {
	ID          int64 `gorm:"primaryKey;autoIncrement:true"`
	IdTrx       int64
	Trx         Trx `gorm:"foreignKey:IdTrx"`
	IdLogProduk int64
	LogProduk   LogProduk `gorm:"foreignKey:IdLogProduk"`
	IdToko      int64
	Kuantitas   int
	HargaTotal  int
	UpdatedAt   time.Time `gorm:"type:date;default:CURRENT_TIMESTAMP()"`
	CreatedAt   time.Time `gorm:"type:date;default:CURRENT_TIMESTAMP()"`
}
