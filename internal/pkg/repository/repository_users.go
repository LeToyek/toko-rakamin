package repository

import (
	"context"
	"fmt"
	"rakamin-final/internal/daos"
	"rakamin-final/internal/helper"
	"reflect"
	"strings"

	"gorm.io/gorm"
)

type UsersRepository interface {
	GetAllUsers(ctx context.Context, params daos.FilterUser) (res []daos.User, err error)
	GetUserByID(ctx context.Context, id int64) (res daos.User, err error)
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
	db := r.db

	structType := reflect.TypeOf(params)
	structValue := reflect.ValueOf(params)

	whereConditions := make([]string, 0)
	whereValues := make([]interface{}, 0)

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		value := structValue.Field(i)

		if value.Interface() != reflect.Zero(field.Type).Interface() {
			whereConditions = append(whereConditions, fmt.Sprintf("%v like ?", field.Name))
			whereValues = append(whereValues, value.Interface())
		}
	}
	if len(whereConditions) > 0 {
		query := strings.Join(whereConditions, " OR ")
		err := db.Where(query, whereValues...).WithContext(ctx).Find(&res).Error
		if err != nil {
			return res, err
		}
	}

	return res, nil
}

func (r *userRepositoryImpl) GetUserByID(ctx context.Context, id int64) (res daos.User, err error) {
	if err := r.db.Where(&daos.User{
		ID: id,
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
