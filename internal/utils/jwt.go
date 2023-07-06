package utils

import (
	"fmt"
	"rakamin-final/internal/helper"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

type (
	JwtCustomClaims struct {
		ID string `json:"id"`
		jwt.StandardClaims
	}
)

func CreateJWT(id string) (res string, err error) {
	currentfilepath := "internal/utils/jwt.go"
	now := time.Now().UTC()
	claims := &JwtCustomClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: now.Add(time.Hour * 24).Unix(),
			IssuedAt:  now.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	fmt.Println(viper.GetString("jwt_secret"))
	tokenString, errToken := token.SignedString([]byte(viper.GetString("jwt_secret")))
	if errToken != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("error when get credential user login, err: %v", errToken.Error()))
		return res, errToken
	}

	return tokenString, nil
}
