package usecase

import (
	"context"
	"errors"
	"fmt"
	"rakamin-final/internal/daos"
	"rakamin-final/internal/helper"
	userdto "rakamin-final/internal/pkg/dto"
	userRepo "rakamin-final/internal/pkg/repository"
	"rakamin-final/internal/utils"
	hasher "rakamin-final/internal/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var currentfilepath = "internal/pkg/usecase/usecase_users.go"

type UsersUsecase interface {
	GetAllUsers(ctx context.Context, params userdto.UserFilter) (res []userdto.UserResponse, err *helper.ErrorStruct)
	GetUserByID(ctx context.Context, id int64) (res userdto.UserResponse, err *helper.ErrorStruct)
	GetCredentialUserLogin(ctx context.Context, params userdto.UserLogin) (res userdto.UserResponse, err *helper.ErrorStruct, token string)
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

func (u *usersUsecaseImpl) GetCredentialUserLogin(ctx context.Context, params userdto.UserLogin) (res userdto.UserResponse, err *helper.ErrorStruct, token string) {

	if err := helper.Validate.Struct(params); err != nil {
		return res, &helper.ErrorStruct{
			Code:    fiber.StatusBadRequest,
			Message: err,
		}, token
	}
	helper.Logger(currentfilepath, helper.LoggerLevelInfo, fmt.Sprintf("params: %v", params))

	resRepo, errRepo := u.repo.GetAllUsers(ctx, daos.FilterUser{
		Email: params.Email,
	})
	if len(resRepo) == 0 {
		return res, &helper.ErrorStruct{
			Code:    fiber.StatusBadRequest,
			Message: errors.New("email not found"),
		}, token
	}

	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code:    fiber.StatusNotFound,
			Message: errors.New("data not found"),
		}, token
	}

	if errRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("error when get credential user login, err: %v", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code:    fiber.StatusInternalServerError,
			Message: errRepo,
		}, token
	}
	if err := utils.VerifyPassword(resRepo[0].KataSandi, params.Password); err != nil {
		return res, &helper.ErrorStruct{
			Code:    fiber.StatusBadRequest,
			Message: errors.New("password not match"),
		}, token
	}

	res = userdto.UserResponse{
		ID:           resRepo[0].ID,
		Nama:         resRepo[0].Nama,
		Email:        resRepo[0].Email,
		NoTelp:       resRepo[0].NoTelp,
		JenisKelamin: resRepo[0].JenisKelamin,
		Pekerjaan:    resRepo[0].Pekerjaan,
		IsAdmin:      resRepo[0].IsAdmin,
	}
	tokenString, errToken := utils.CreateJWT(strconv.FormatInt(res.ID, 10), res.IsAdmin)
	if errToken != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("error when create token, err: %v", errToken.Error()))
		return res, &helper.ErrorStruct{
			Code:    fiber.StatusInternalServerError,
			Message: errToken,
		}, token
	}
	token = tokenString

	return res, nil, token
}

func (u *usersUsecaseImpl) GetUserByID(ctx context.Context, id int64) (res userdto.UserResponse, err *helper.ErrorStruct) {
	resRepo, errRepo := u.repo.GetUserByID(ctx, id)

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
	if errValidate := helper.Validate.Struct(user); errValidate != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("error when validate struct, err: %v", errValidate.Error()))
		return res, &helper.ErrorStruct{
			Code:    fiber.StatusBadRequest,
			Message: errValidate,
		}
	}
	resHash, errHash := hasher.HashPassword(user.Password)
	if errHash != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("error when hash password, err: %v", errHash.Error()))
		return res, &helper.ErrorStruct{
			Code:    fiber.StatusInternalServerError,
			Message: errHash,
		}
	}

	resRepo, errRepo := u.repo.CreateUser(ctx, daos.User{
		Nama:         user.Nama,
		KataSandi:    resHash,
		IdProvinsi:   user.IdProvinsi,
		IdKota:       user.IdKota,
		Email:        user.Email,
		Tentang:      user.Tentang,
		NoTelp:       user.NoTelp,
		JenisKelamin: user.JenisKelamin,
		Pekerjaan:    user.Pekerjaan,
		IsAdmin:      user.IsAdmin,
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
		IsAdmin:      resRepo.IsAdmin,
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
