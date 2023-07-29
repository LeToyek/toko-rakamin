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
		KodeInvoice string `gorm:"size:255;index:idx_kodeinvoice,unique"`
		MethodBayar string `gorm:"size:255"`
		IdToko      int64
		Toko        Toko        `gorm:"foreignKey:IdToko"`
		DetailTrxes []DetailTrx `gorm:"foreignKey:IdTrx"`
		UpdatedAt   time.Time   `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP()"`
		CreatedAt   time.Time   `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP()"`
	}

	FilterTrx struct {
		ID            int64
		Limit, Offset int
		KodeInvoice   string
	}
)
