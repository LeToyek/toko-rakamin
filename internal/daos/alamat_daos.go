package daos

import (
	"time"

	"gorm.io/gorm"
)

type (
	Alamat struct {
		gorm.Model
		ID           int64 `gorm:"primaryKey;autoIncrement:true"`
		IdUser       int64
		User         User      `gorm:"foreignKey:IdUser"`
		JudulAlamat  string    `gorm:"size:255"`
		NamaPenerima string    `gorm:"size:255"`
		NoTelp       string    `gorm:"size:255"`
		DetailAlamat string    `gorm:"size:255"`
		UpdatedAt    time.Time `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP()"`
		CreatedAt    time.Time `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP()"`
	}

	FilterAlamat struct {
		ID            int64
		Limit, Offset int
		JudulAlamat   string
	}
)
