package middleware

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

func DeserializeUser(c *fiber.Ctx) error {
	var tokenString string

	auth := c.Get("Authorization")

	if strings.HasPrefix(auth, "Bearer ") {
		tokenString = strings.TrimPrefix(auth, "Bearer ")
	} else if c.Cookies("token") != "" {
		tokenString = c.Cookies("token")
	}
	fmt.Print("tokenString: ", c.Cookies("token"))
	fmt.Println(tokenString)

	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	tokenByte, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.ErrUnauthorized
		}

		return []byte(viper.GetString("jwt_secret")), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	claims, ok := tokenByte.Claims.(jwt.MapClaims)

	if !ok || !tokenByte.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	fmt.Println(claims)
	return c.Next()
}
