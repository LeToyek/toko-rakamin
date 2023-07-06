package daos

import (
	"time"
)

type (
	Trx struct {
		ID          int64 `gorm:"primaryKey;autoIncrement:true"`
		IdUser      int64
		User        User `gorm:"foreignKey:IdUser"`
		IdAlamat    int64
		Alamat      Alamat `gorm:"foreignKey:IdAlamat"`
		HargaTotal  int64
		KodeInvoice string    `gorm:"size:255"`
		MethodBayar string    `gorm:"size:255"`
		UpdatedAt   time.Time `gorm:"type:date;default:CURRENT_TIMESTAMP()"`
		CreatedAt   time.Time `gorm:"type:date;default:CURRENT_TIMESTAMP()"`
	}
)
