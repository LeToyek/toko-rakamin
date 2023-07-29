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
		IdAlamat    int64              `json:"id_alamat" validate:"required"`
		MethodBayar string             `json:"method_bayar" validate:"required"`
		DetailTrxes []DetailTrxRequest `json:"detail_trxes" validate:"required"`
	}
	DetailTrxRequest struct {
		IdProduk  int64 `json:"id_produk" validate:"required"`
		Kuantitas int64 `json:"kuantitas" validate:"required"`
	}
	DetailTrxResponse struct {
		ProductID int64 `json:"produk_id"`
		Kuantitas int64 `json:"kuantitas"`
	}
	TrxResponse struct {
		ID          int64               `json:"id"`
		IdUser      int64               `json:"id_user"`
		HargaTotal  int64               `json:"harga_total"`
		KodeInvoice string              `json:"kode_invoice"`
		MethodBayar string              `json:"method_bayar"`
		Alamat      AddressResponse     `json:"alamat_kirim"`
		Toko        StoreResponse       `json:"toko"`
		DetailTrxes []DetailTrxResponse `json:"detail_trxes"`
		UpdatedAt   time.Time           `json:"updated_at"`
		CreatedAt   time.Time           `json:"created_at"`
	}
)
