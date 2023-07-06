package repository

import (
	"context"
	"fmt"
	"rakamin-final/internal/daos"
	"rakamin-final/internal/helper"

	"gorm.io/gorm"
)

type UsersRepository interface {
	GetAllUsers(ctx context.Context, params daos.FilterUser) (res []daos.User, err error)
	GetUserByParam(ctx context.Context, params daos.FilterUser) (res daos.User, err error)
	CreateUser(ctx context.Context, user daos.User) (res daos.User, err error)
	UpdateUser(ctx context.Context, id int64, user daos.User) (res daos.User, err error)
	DeleteUser(ctx context.Context, id int64) error
}

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepositoryImpl {
	return &userRepositoryImpl{db}
}

func (r *userRepositoryImpl) GetAllUsers(ctx context.Context, params daos.FilterUser) (res []daos.User, err error) {
	return res, nil
}

func (r *userRepositoryImpl) GetUserByParam(ctx context.Context, params daos.FilterUser) (res daos.User, err error) {
	if err := r.db.Where(&daos.User{
		Email:     params.Email,
		KataSandi: params.Password,
	}).First(&res).WithContext(ctx).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (r *userRepositoryImpl) CreateUser(ctx context.Context, user daos.User) (res daos.User, err error) {
	if err := r.db.Create(&user).WithContext(ctx).Error; err != nil {
		return res, err
	}
	helper.Logger("internal/pkg/repository/repository_users.go", helper.LoggerLevelInfo, fmt.Sprintf("user created, id: %v", user))
	return user, nil
}

func (r *userRepositoryImpl) UpdateUser(ctx context.Context, id int64, user daos.User) (res daos.User, err error) {
	if err := r.db.Where(&daos.User{
		ID: id,
	}).Updates(&user).WithContext(ctx).Error; err != nil {
		return res, err
	}
	return user, nil
}

func (r *userRepositoryImpl) DeleteUser(ctx context.Context, id int64) error {
	if err := r.db.Delete(&daos.User{}, id).WithContext(ctx).Error; err != nil {
		return err
	}
	return nil
}
