package usecase

import (
	"context"
	"rakamin-final/internal/daos"
	"rakamin-final/internal/helper"
	dto "rakamin-final/internal/pkg/dto"
	"rakamin-final/internal/pkg/repository"
)

type TrxUsecase interface {
	GetAllTrxes(ctx context.Context, params dto.FilterTrx) (res []dto.TrxResponse, err *helper.ErrorStruct)
	GetTrxByID(ctx context.Context, id int64) (res dto.TrxResponse, err *helper.ErrorStruct)
	CreateTrx(ctx context.Context, TrxData dto.TrxRequest) (res dto.TrxResponse, err *helper.ErrorStruct)
	UpdateTrx(ctx context.Context, id int64, TrxData dto.TrxRequest) (res dto.TrxResponse, err *helper.ErrorStruct)
	DeleteTrx(ctx context.Context, id int64) *helper.ErrorStruct
}

type trxUsecaseImpl struct {
	trxRepo repository.TrxRepository
}

func NewTrxUsecase(trxRepo repository.TrxRepository) *trxUsecaseImpl {
	return &trxUsecaseImpl{trxRepo}
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
		res = append(res, dto.TrxResponse{
			ID:          v.ID,
			KodeInvoice: v.KodeInvoice,
			HargaTotal:  v.HargaTotal,
			MethodBayar: v.MethodBayar,
			IdUser:      v.IdUser,
			IdAlamat:    v.IdAlamat,
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
	res = dto.TrxResponse{
		ID:          resRepo.ID,
		KodeInvoice: resRepo.KodeInvoice,
		HargaTotal:  resRepo.HargaTotal,
		MethodBayar: resRepo.MethodBayar,
		IdUser:      resRepo.IdUser,
		IdAlamat:    resRepo.IdAlamat,
		CreatedAt:   resRepo.CreatedAt,
		UpdatedAt:   resRepo.UpdatedAt,
	}
	return res, nil
}

func (u *trxUsecaseImpl) CreateTrx(ctx context.Context, TrxData dto.TrxRequest) (res dto.TrxResponse, err *helper.ErrorStruct) {
	resRepo, errRepo := u.trxRepo.CreateTrx(ctx, daos.Trx{
		KodeInvoice: TrxData.KodeInvoice,
		HargaTotal:  TrxData.HargaTotal,
		MethodBayar: TrxData.MethodBayar,
		IdUser:      TrxData.IdUser,
		IdAlamat:    TrxData.IdAlamat,
	})
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
		IdAlamat:    resRepo.IdAlamat,
		CreatedAt:   resRepo.CreatedAt,
		UpdatedAt:   resRepo.UpdatedAt,
	}
	return res, nil
}

func (u *trxUsecaseImpl) UpdateTrx(ctx context.Context, id int64, TrxData dto.TrxRequest) (res dto.TrxResponse, err *helper.ErrorStruct) {
	resRepo, errRepo := u.trxRepo.UpdateTrx(ctx, id, daos.Trx{
		KodeInvoice: TrxData.KodeInvoice,
		HargaTotal:  TrxData.HargaTotal,
		MethodBayar: TrxData.MethodBayar,
		IdUser:      TrxData.IdUser,
		IdAlamat:    TrxData.IdAlamat,
	})
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
		IdAlamat:    resRepo.IdAlamat,
		CreatedAt:   resRepo.CreatedAt,
		UpdatedAt:   resRepo.UpdatedAt,
	}
	return res, nil
}

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
