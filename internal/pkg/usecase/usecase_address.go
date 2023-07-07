package usecase

import (
	"context"
	"rakamin-final/internal/daos"
	"rakamin-final/internal/helper"
	dto "rakamin-final/internal/pkg/dto"
	"rakamin-final/internal/pkg/repository"
)

type AddressUsecase interface {
	CreateAddress(c context.Context, address dto.AddressRequest) (res dto.AddressResponse, err *helper.ErrorStruct)
	GetAllAddress(c context.Context, params dto.AddressFilter) (res []dto.AddressResponse, err *helper.ErrorStruct)
	GetAddressByID(c context.Context, id int64) (res dto.AddressResponse, err *helper.ErrorStruct)
	UpdateAddress(c context.Context, id int64, address dto.AddressRequest) (res dto.AddressResponse, err *helper.ErrorStruct)
	DeleteAddress(c context.Context, id int64) *helper.ErrorStruct
}

type addressUsecaseImpl struct {
	addressRepo repository.AddressRepository
}

func NewAddressUsecase(addressRepo repository.AddressRepository) *addressUsecaseImpl {
	return &addressUsecaseImpl{addressRepo}
}

func (u *addressUsecaseImpl) CreateAddress(c context.Context, address dto.AddressRequest) (res dto.AddressResponse, err *helper.ErrorStruct) {
	resRepo, errRepo := u.addressRepo.CreateAddress(c, daos.Alamat{
		JudulAlamat:  address.JudulAlamat,
		NamaPenerima: address.NamaPenerima,
		NoTelp:       address.NoTelp,
		DetailAlamat: address.DetailAlamat,
	})

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Code:    500,
			Message: errRepo,
		}
	}

	res = dto.AddressResponse{
		ID:           resRepo.ID,
		JudulAlamat:  resRepo.JudulAlamat,
		NamaPenerima: resRepo.NamaPenerima,
		NoTelp:       resRepo.NoTelp,
		DetailAlamat: resRepo.DetailAlamat,
	}

	return res, nil
}

func (u *addressUsecaseImpl) GetAllAddress(c context.Context, params dto.AddressFilter) (res []dto.AddressResponse, err *helper.ErrorStruct) {
	if params.Limit < 1 {
		params.Limit = 10
	}
	if params.Page < 1 {
		params.Page = 0
	} else {
		params.Page = (params.Page - 1) * params.Limit
	}
	resRepo, errRepo := u.addressRepo.GetAllAddress(c, daos.FilterAlamat{
		ID:          int64(params.ID),
		Limit:       params.Limit,
		JudulAlamat: params.Name,
	})

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Code:    500,
			Message: errRepo,
		}
	}
	for _, v := range resRepo {
		res = append(res, dto.AddressResponse{
			ID:           v.ID,
			JudulAlamat:  v.JudulAlamat,
			NamaPenerima: v.NamaPenerima,
			NoTelp:       v.NoTelp,
			DetailAlamat: v.DetailAlamat,
		})
	}

	return res, nil
}

func (u *addressUsecaseImpl) GetAddressByID(c context.Context, id int64) (res dto.AddressResponse, err *helper.ErrorStruct) {
	resRepo, errRepo := u.addressRepo.GetAddressByID(c, id)

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Code:    500,
			Message: errRepo,
		}
	}

	res = dto.AddressResponse{
		ID:           resRepo.ID,
		JudulAlamat:  resRepo.JudulAlamat,
		NamaPenerima: resRepo.NamaPenerima,
		NoTelp:       resRepo.NoTelp,
		DetailAlamat: resRepo.DetailAlamat,
	}

	return res, nil
}

func (u *addressUsecaseImpl) UpdateAddress(c context.Context, id int64, address dto.AddressRequest) (res dto.AddressResponse, err *helper.ErrorStruct) {
	resRepo, errRepo := u.addressRepo.UpdateAddress(c, id, daos.Alamat{
		JudulAlamat:  address.JudulAlamat,
		NamaPenerima: address.NamaPenerima,
		NoTelp:       address.NoTelp,
		DetailAlamat: address.DetailAlamat,
	})

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Code:    500,
			Message: errRepo,
		}
	}

	res = dto.AddressResponse{
		ID:           resRepo.ID,
		JudulAlamat:  resRepo.JudulAlamat,
		NamaPenerima: resRepo.NamaPenerima,
		NoTelp:       resRepo.NoTelp,
		DetailAlamat: resRepo.DetailAlamat,
	}

	return res, nil
}

func (u *addressUsecaseImpl) DeleteAddress(c context.Context, id int64) *helper.ErrorStruct {
	errRepo := u.addressRepo.DeleteAddress(c, id)

	if errRepo != nil {
		return &helper.ErrorStruct{
			Code:    500,
			Message: errRepo,
		}
	}

	return nil
}
