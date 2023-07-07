package dto

import (
	"time"
)

type (
	Category struct {
		ID        int64     `json:"id"`
		Nama      string    `json:"nama"`
		UpdatedAt time.Time `json:"updated_at"`
		CreatedAt time.Time `json:"created_at"`
	}

	CategoryRequest struct {
		Nama string `json:"nama"`
	}

	CategoryResponse struct {
		ID        int64     `json:"id"`
		Nama      string    `json:"nama"`
		UpdatedAt time.Time `json:"updated_at"`
		CreatedAt time.Time `json:"created_at"`
	}
)
