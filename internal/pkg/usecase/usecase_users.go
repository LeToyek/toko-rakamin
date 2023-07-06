package usecase

import (
	"context"
	"errors"
	"fmt"
	"rakamin-final/internal/daos"
	"rakamin-final/internal/helper"
	userdto "rakamin-final/internal/pkg/dto"
	userRepo "rakamin-final/internal/pkg/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var currentfilepath = "internal/pkg/usecase/usecase_users.go"

type UsersUsecase interface {
	GetAllUsers(ctx context.Context, params userdto.UserFilter) (res []userdto.UserResponse, err *helper.ErrorStruct)
	GetUserByParam(ctx context.Context, params userdto.UserFilter) (res userdto.UserResponse, err *helper.ErrorStruct)
	CreateUser(ctx context.Context, user userdto.UserRegister) (res userdto.UserResponse, err *helper.ErrorStruct)
	UpdateUser(ctx context.Context, id int64, user userdto.UserRegister) (res userdto.UserResponse, err *helper.ErrorStruct)
	DeleteUser(ctx context.Context, id int64) *helper.ErrorStruct
}

type usersUsecaseImpl struct {
	repo userRepo.UsersRepository
}

func NewUsersUsecase(userRepo userRepo.UsersRepository) *usersUsecaseImpl {
	return &usersUsecaseImpl{
		repo: userRepo,
	}
}

func (u *usersUsecaseImpl) GetAllUsers(ctx context.Context, params userdto.UserFilter) (res []userdto.UserResponse, err *helper.ErrorStruct) {
	if params.Limit < 1 {
		params.Limit = 10
	}
	if params.Page < 1 {
		params.Page = 0
	} else {
		params.Page = (params.Page - 1) * params.Limit
	}

	resRepo, errRepo := u.repo.GetAllUsers(ctx, daos.FilterUser{
		Limit:  params.Limit,
		Offset: params.Page,
		Email:  params.Email,
	})

	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code:    fiber.StatusNotFound,
			Message: errors.New("data not found"),
		}
	}

	if errRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("error when get all users, err: %v", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code:    fiber.StatusInternalServerError,
			Message: errRepo,
		}

	}
	for _, v := range resRepo {
		res = append(res, userdto.UserResponse{
			ID:           v.ID,
			Nama:         v.Nama,
			Email:        v.Email,
			NoTelp:       v.NoTelp,
			JenisKelamin: v.JenisKelamin,
			Pekerjaan:    v.Pekerjaan,
		})

	}
	return res, nil
}

func (u *usersUsecaseImpl) GetUserByParam(ctx context.Context, params userdto.UserFilter) (res userdto.UserResponse, err *helper.ErrorStruct) {
	resRepo, errRepo := u.repo.GetUserByParam(ctx, daos.FilterUser{
		Email:    params.Email,
		Password: params.Password,
	})

	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code:    fiber.StatusNotFound,
			Message: errors.New("data not found"),
		}
	}

	if errRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("error when get user by param, err: %v", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code:    fiber.StatusInternalServerError,
			Message: errRepo,
		}

	}

	res = userdto.UserResponse{
		ID:           resRepo.ID,
		Nama:         resRepo.Nama,
		Email:        resRepo.Email,
		NoTelp:       resRepo.NoTelp,
		JenisKelamin: resRepo.JenisKelamin,
		Pekerjaan:    resRepo.Pekerjaan,
	}
	return res, nil
}

func (u *usersUsecaseImpl) CreateUser(ctx context.Context, user userdto.UserRegister) (res userdto.UserResponse, err *helper.ErrorStruct) {
	resRepo, errRepo := u.repo.CreateUser(ctx, daos.User{
		Nama:         user.Nama,
		KataSandi:    user.Password,
		IdProvinsi:   user.IdProvinsi,
		IdKota:       user.IdKota,
		Email:        user.Email,
		Tentang:      user.Tentang,
		NoTelp:       user.NoTelp,
		JenisKelamin: user.JenisKelamin,
		Pekerjaan:    user.Pekerjaan,
	})
	helper.Logger(currentfilepath, helper.LoggerLevelInfo, fmt.Sprintf("resRepo: %v", resRepo))

	if errRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("error when create user, err: %v", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code:    fiber.StatusInternalServerError,
			Message: errRepo,
		}

	}

	res = userdto.UserResponse{
		ID:           resRepo.ID,
		Nama:         resRepo.Nama,
		Email:        resRepo.Email,
		NoTelp:       resRepo.NoTelp,
		JenisKelamin: resRepo.JenisKelamin,
		Pekerjaan:    resRepo.Pekerjaan,
	}
	return res, nil
}

func (u *usersUsecaseImpl) UpdateUser(ctx context.Context, id int64, user userdto.UserRegister) (res userdto.UserResponse, err *helper.ErrorStruct) {
	resRepo, errRepo := u.repo.UpdateUser(ctx, id, daos.User{
		Nama:         user.Nama,
		Email:        user.Email,
		NoTelp:       user.NoTelp,
		JenisKelamin: user.JenisKelamin,
		Pekerjaan:    user.Pekerjaan,
	})

	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code:    fiber.StatusNotFound,
			Message: errors.New("data not found"),
		}
	}

	if errRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("error when update user, err: %v", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code:    fiber.StatusInternalServerError,
			Message: errRepo,
		}

	}

	res = userdto.UserResponse{
		ID:           resRepo.ID,
		Nama:         resRepo.Nama,
		Email:        resRepo.Email,
		NoTelp:       resRepo.NoTelp,
		JenisKelamin: resRepo.JenisKelamin,
		Pekerjaan:    resRepo.Pekerjaan,
	}
	return res, nil
}

func (u *usersUsecaseImpl) DeleteUser(ctx context.Context, id int64) (err *helper.ErrorStruct) {
	errRepo := u.repo.DeleteUser(ctx, id)

	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return &helper.ErrorStruct{
			Code:    fiber.StatusNotFound,
			Message: errors.New("data not found"),
		}
	}

	if errRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("error when delete user, err: %v", errRepo.Error()))
		return &helper.ErrorStruct{
			Code:    fiber.StatusInternalServerError,
			Message: errRepo,
		}

	}

	return nil
}
