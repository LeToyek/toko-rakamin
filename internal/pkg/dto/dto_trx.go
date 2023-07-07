package dto

import (
	"time"
)

type (
	FilterTrx struct {
		ID            int64
		Limit, Offset int
		KodeInvoice   string
	}

	TrxRequest struct {
		IdUser      int64  `json:"id_user" validate:"required"`
		IdAlamat    int64  `json:"id_alamat" validate:"required"`
		HargaTotal  int64  `json:"harga_total" validate:"required"`
		KodeInvoice string `json:"kode_invoice" validate:"required"`
		MethodBayar string `json:"method_bayar" validate:"required"`
	}

	TrxResponse struct {
		ID          int64     `json:"id"`
		IdUser      int64     `json:"id_user"`
		IdAlamat    int64     `json:"id_alamat"`
		HargaTotal  int64     `json:"harga_total"`
		KodeInvoice string    `json:"kode_invoice"`
		MethodBayar string    `json:"method_bayar"`
		UpdatedAt   time.Time `json:"updated_at"`
		CreatedAt   time.Time `json:"created_at"`
	}
)
