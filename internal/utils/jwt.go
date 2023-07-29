package utils

import (
	"fmt"
	"rakamin-final/internal/helper"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

type (
	JwtCustomClaims struct {
		ID      string `json:"id"`
		IsAdmin bool   `json:"is_admin"`
		jwt.StandardClaims
	}
)

func CreateJWT(id string, isAdmin bool) (res string, err error) {
	currentfilepath := "internal/utils/jwt.go"
	now := time.Now().UTC()
	claims := &JwtCustomClaims{
		id,
		isAdmin,
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

func GetJWTCredentials(tokenString string, c *fiber.Ctx) (*JwtCustomClaims, error) {
	tokenByte, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.ErrUnauthorized
		}

		return []byte(viper.GetString("jwt_secret")), nil
	})

	if err != nil {
		return nil, c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	claims, ok := tokenByte.Claims.(jwt.MapClaims)

	if !ok || !tokenByte.Valid {
		return nil, c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	return &JwtCustomClaims{
		ID:      claims["id"].(string),
		IsAdmin: claims["is_admin"].(bool),
	}, nil
}
