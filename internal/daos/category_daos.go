package daos

import (
	"time"
)

type (
	Category struct {
		ID           int64     `gorm:"primaryKey;autoIncrement:true"`
		NamaCategory string    `gorm:"size:255"`
		UpdatedAt    time.Time `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP()"`
		CreatedAt    time.Time `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP()"`
	}
)
