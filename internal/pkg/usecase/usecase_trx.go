package usecase

import (
	"context"
	"rakamin-final/internal/daos"
	"rakamin-final/internal/helper"
	dto "rakamin-final/internal/pkg/dto"
	"rakamin-final/internal/pkg/repository"
	"rakamin-final/internal/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type TrxUsecase interface {
	GetAllTrxes(ctx context.Context, params dto.FilterTrx) (res []dto.TrxResponse, err *helper.ErrorStruct)
	GetTrxByID(ctx context.Context, id int64) (res dto.TrxResponse, err *helper.ErrorStruct)
	CreateTrx(ctx context.Context, TrxData dto.TrxRequest, userID int64) (res dto.TrxResponse, err *helper.ErrorStruct)
	// UpdateTrx(ctx context.Context, id int64, userID int64, TrxData dto.TrxRequest) (res dto.TrxResponse, err *helper.ErrorStruct)
	DeleteTrx(ctx context.Context, id int64) *helper.ErrorStruct
}

type trxUsecaseImpl struct {
	trxRepo       repository.TrxRepository
	detailTrxRepo repository.DetailTrxRepository
	productRepo   repository.ProductRepository
	addressRepo   repository.AddressRepository
}

func NewTrxUsecase(
	trxRepo repository.TrxRepository,
	detailTrxRepo repository.DetailTrxRepository,
	productRepo repository.ProductRepository,
	addressRepo repository.AddressRepository) *trxUsecaseImpl {
	return &trxUsecaseImpl{trxRepo, detailTrxRepo, productRepo, addressRepo}
}

func (u *trxUsecaseImpl) GetAllTrxes(ctx context.Context, params dto.FilterTrx) (res []dto.TrxResponse, err *helper.ErrorStruct) {
	resRepo, errRepo := u.trxRepo.GetAllTrxes(ctx, daos.FilterTrx{
		ID:          int64(params.ID),
		Limit:       params.Limit,
		Offset:      params.Offset,
		KodeInvoice: params.KodeInvoice,
	})
	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Code:    500,
			Message: errRepo,
		}
	}

	for _, v := range resRepo {
		respAdress := dto.AddressResponse{
			ID:           v.Alamat.ID,
			JudulAlamat:  v.Alamat.JudulAlamat,
			NamaPenerima: v.Alamat.NamaPenerima,
			NoTelp:       v.Alamat.NoTelp,
			DetailAlamat: v.Alamat.DetailAlamat,
			User: dto.UserResponse{
				ID:           v.Alamat.User.ID,
				Nama:         v.Alamat.User.Nama,
				Email:        v.Alamat.User.Email,
				NoTelp:       v.Alamat.User.NoTelp,
				JenisKelamin: v.Alamat.User.JenisKelamin,
				Pekerjaan:    v.Alamat.User.Pekerjaan,
				IsAdmin:      v.Alamat.User.IsAdmin,
			},
		}
		var arrDetailTrxes []dto.DetailTrxResponse
		for _, x := range v.DetailTrxes {
			arrDetailTrxes = append(arrDetailTrxes, dto.DetailTrxResponse{
				ProductID: x.IdProduk,
				Kuantitas: int64(x.Kuantitas),
			})
		}
		resStore := dto.StoreResponse{
			ID:        v.Toko.ID,
			NamaToko:  v.Toko.NamaToko,
			UrlFoto:   v.Toko.UrlFoto,
			UpdatedAt: v.Toko.UpdatedAt,
			CreatedAt: v.Toko.CreatedAt,
			User: dto.UserResponse{
				ID:           v.Toko.User.ID,
				Nama:         v.Toko.User.Nama,
				Email:        v.Toko.User.Email,
				NoTelp:       v.Toko.User.NoTelp,
				JenisKelamin: v.Toko.User.JenisKelamin,
				Pekerjaan:    v.Toko.User.Pekerjaan,
				IsAdmin:      v.Toko.User.IsAdmin,
			},
		}

		res = append(res, dto.TrxResponse{
			ID:          v.ID,
			KodeInvoice: v.KodeInvoice,
			HargaTotal:  v.HargaTotal,
			MethodBayar: v.MethodBayar,
			IdUser:      v.IdUser,
			Alamat:      respAdress,
			DetailTrxes: arrDetailTrxes,
			Toko:        resStore,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		})
	}
	return res, nil
}

func (u *trxUsecaseImpl) GetTrxByID(ctx context.Context, id int64) (res dto.TrxResponse, err *helper.ErrorStruct) {
	resRepo, errRepo := u.trxRepo.GetTrxByID(ctx, id)
	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Code:    500,
			Message: errRepo,
		}
	}
	var arrDetailTrxes []dto.DetailTrxResponse
	for _, v := range resRepo.DetailTrxes {
		arrDetailTrxes = append(arrDetailTrxes, dto.DetailTrxResponse{
			ProductID: v.IdProduk,
			Kuantitas: int64(v.Kuantitas),
		})
	}
	respAdress := dto.AddressResponse{
		ID:           resRepo.Alamat.ID,
		JudulAlamat:  resRepo.Alamat.JudulAlamat,
		NamaPenerima: resRepo.Alamat.NamaPenerima,
		NoTelp:       resRepo.Alamat.NoTelp,
		DetailAlamat: resRepo.Alamat.DetailAlamat,
		User: dto.UserResponse{
			ID:           resRepo.Alamat.User.ID,
			Nama:         resRepo.Alamat.User.Nama,
			Email:        resRepo.Alamat.User.Email,
			NoTelp:       resRepo.Alamat.User.NoTelp,
			JenisKelamin: resRepo.Alamat.User.JenisKelamin,
			Pekerjaan:    resRepo.Alamat.User.Pekerjaan,
			IsAdmin:      resRepo.Alamat.User.IsAdmin,
		},
	}
	resStore := dto.StoreResponse{
		ID:        resRepo.Toko.ID,
		NamaToko:  resRepo.Toko.NamaToko,
		UrlFoto:   resRepo.Toko.UrlFoto,
		UpdatedAt: resRepo.Toko.UpdatedAt,
		CreatedAt: resRepo.Toko.CreatedAt,
		User: dto.UserResponse{
			ID:           resRepo.Toko.User.ID,
			Nama:         resRepo.Toko.User.Nama,
			Email:        resRepo.Toko.User.Email,
			NoTelp:       resRepo.Toko.User.NoTelp,
			JenisKelamin: resRepo.Toko.User.JenisKelamin,
			Pekerjaan:    resRepo.Toko.User.Pekerjaan,
			IsAdmin:      resRepo.Toko.User.IsAdmin,
		},
	}

	res = dto.TrxResponse{
		ID:          resRepo.ID,
		KodeInvoice: resRepo.KodeInvoice,
		HargaTotal:  resRepo.HargaTotal,
		MethodBayar: resRepo.MethodBayar,
		IdUser:      resRepo.IdUser,
		Alamat:      respAdress,
		Toko:        resStore,
		DetailTrxes: arrDetailTrxes,
		CreatedAt:   resRepo.CreatedAt,
		UpdatedAt:   resRepo.UpdatedAt,
	}
	return res, nil
}

func (u *trxUsecaseImpl) CreateTrx(ctx context.Context, TrxData dto.TrxRequest, userID int64) (res dto.TrxResponse, err *helper.ErrorStruct) {
	if err := helper.Validate.Struct(TrxData); err != nil {
		return res, &helper.ErrorStruct{
			Code:    fiber.StatusBadRequest,
			Message: err,
		}
	}

	var arrDetailTrxes []dto.DetailTrxResponse
	var arrDaosDetailTrxes []daos.DetailTrx
	totalPrice := 0

	for _, v := range TrxData.DetailTrxes {
		resProduct, err := u.productRepo.GetProductByID(ctx, v.IdProduk)
		if err != nil {
			return res, &helper.ErrorStruct{
				Code:    500,
				Message: err,
			}
		}
		productPrice, err := strconv.Atoi(resProduct.HargaKonsumen)
		if err != nil {
			return res, &helper.ErrorStruct{
				Code:    500,
				Message: err,
			}
		}
		totalDetailPrice := v.Kuantitas * int64(productPrice)

		newDaosDetailTrx := daos.DetailTrx{
			IdProduk:   resProduct.ID,
			Kuantitas:  int(v.Kuantitas),
			HargaTotal: int(totalDetailPrice),
			IdToko:     resProduct.IdToko,
		}
		arrDaosDetailTrxes = append(arrDaosDetailTrxes, newDaosDetailTrx)

		newDetailTrx := dto.DetailTrxResponse{
			ProductID: resProduct.ID,
			Kuantitas: v.Kuantitas,
		}
		arrDetailTrxes = append(arrDetailTrxes, newDetailTrx)
		totalPrice += int(totalDetailPrice)
	}
	resAddressRepo, errAdressRepo := u.addressRepo.GetAddressByID(ctx, TrxData.IdAlamat)
	if errAdressRepo != nil {
		return res, &helper.ErrorStruct{
			Code:    500,
			Message: errAdressRepo,
		}
	}
	respAdress := dto.AddressResponse{
		ID:           resAddressRepo.ID,
		JudulAlamat:  resAddressRepo.JudulAlamat,
		NamaPenerima: resAddressRepo.NamaPenerima,
		NoTelp:       resAddressRepo.NoTelp,
		DetailAlamat: resAddressRepo.DetailAlamat}

	idToko := userID
	invoiceCode := utils.GenerateInvoiceCode()
	resRepo, errRepo := u.trxRepo.CreateTrx(ctx, daos.Trx{
		KodeInvoice: invoiceCode,
		HargaTotal:  int64(totalPrice),
		MethodBayar: TrxData.MethodBayar,
		IdUser:      userID,
		IdAlamat:    TrxData.IdAlamat,
		IdToko:      idToko,
	}, arrDaosDetailTrxes)
	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Code:    500,
			Message: errRepo,
		}
	}

	res = dto.TrxResponse{
		ID:          resRepo.ID,
		KodeInvoice: resRepo.KodeInvoice,
		HargaTotal:  resRepo.HargaTotal,
		MethodBayar: resRepo.MethodBayar,
		IdUser:      resRepo.IdUser,
		Alamat:      respAdress,
		CreatedAt:   resRepo.CreatedAt,
		UpdatedAt:   resRepo.UpdatedAt,
		DetailTrxes: arrDetailTrxes,
	}
	return res, nil
}

// func (u *trxUsecaseImpl) UpdateTrx(ctx context.Context, id int64, userID int64, TrxData dto.TrxRequest) (res dto.TrxResponse, err *helper.ErrorStruct) {
// 	resRepo, errRepo := u.trxRepo.UpdateTrx(ctx, id, daos.Trx{
// 		KodeInvoice: TrxData.KodeInvoice,
// 		HargaTotal:  TrxData.HargaTotal,
// 		MethodBayar: TrxData.MethodBayar,
// 		IdUser:      userID,
// 		IdAlamat:    TrxData.IdAlamat,
// 	})
// 	if errRepo != nil {
// 		return res, &helper.ErrorStruct{
// 			Code:    500,
// 			Message: errRepo,
// 		}
// 	}
// 	resAddressRepo, errAdressRepo := u.addressRepo.GetAddressByID(ctx, resRepo.IdAlamat)
// 	if errAdressRepo != nil {
// 		return res, &helper.ErrorStruct{
// 			Code:    500,
// 			Message: errAdressRepo,
// 		}
// 	}
// 	respAdress := dto.AddressResponse{
// 		ID:           resAddressRepo.ID,
// 		JudulAlamat:  resAddressRepo.JudulAlamat,
// 		NamaPenerima: resAddressRepo.NamaPenerima,
// 		NoTelp:       resAddressRepo.NoTelp,
// 		DetailAlamat: resAddressRepo.DetailAlamat}

// 	res = dto.TrxResponse{
// 		ID:          resRepo.ID,
// 		KodeInvoice: resRepo.KodeInvoice,
// 		HargaTotal:  resRepo.HargaTotal,
// 		MethodBayar: resRepo.MethodBayar,
// 		IdUser:      resRepo.IdUser,
// 		Alamat:      respAdress,
// 		CreatedAt:   resRepo.CreatedAt,
// 		UpdatedAt:   resRepo.UpdatedAt,
// 	}
// 	return res, nil
// }

func (u *trxUsecaseImpl) DeleteTrx(ctx context.Context, id int64) *helper.ErrorStruct {
	errRepo := u.trxRepo.DeleteTrx(ctx, id)
	if errRepo != nil {
		return &helper.ErrorStruct{
			Code:    500,
			Message: errRepo,
		}
	}
	return nil
}
