package daos

import (
	"time"

	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model
		ID           int64  `gorm:"primaryKey;autoIncrement:true"`
		Nama         string `gorm:"size:255"`
		KataSandi    string `gorm:"size:255"`
		NoTelp       string `gorm:"size:255;index:idx_notelp,unique"`
		JenisKelamin string `gorm:"size:255"`
		Email        string `gorm:"size:255;index:idx_email,unique"`
		Tentang      string `gorm:"type:text"`
		Pekerjaan    string `gorm:"size:255"`
		IdProvinsi   string `gorm:"size:255"`
		IdKota       string `gorm:"size:255"`
		IsAdmin      bool
		Updated_at   time.Time `gorm:"type:date;default:CURRENT_TIMESTAMP()"`
		Created_at   time.Time `gorm:"type:date;default:CURRENT_TIMESTAMP()"`
		Alamats      []Alamat  `gorm:"foreignKey:IdUser"`
	}

	FilterUser struct {
		ID              int64
		Limit, Offset   int
		Email, Password string
	}
)
