package middleware

import (
	"fmt"
	"strings"

	"github.com/forceki/invent-be/config"
	"github.com/forceki/invent-be/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Auth(f *fiber.Ctx) error {
	var tokenString string
	authorization := f.Get("Authorization")

	if strings.HasPrefix(authorization, "Bearer ") {
		tokenString = strings.TrimPrefix(authorization, "Bearer ")
	} else if f.Cookies("token") != "" {
		tokenString = f.Cookies("token")
	}

	if tokenString == "" {
		return handler.ResponseHttp(f, fiber.StatusUnauthorized, 0, "your not have permission", nil)
	}

	JwtSecret := config.Config("JWT_SECRET")

	tokenByte, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}

		return []byte(JwtSecret), nil
	})

	if err != nil {
		return handler.ResponseHttp(f, fiber.StatusUnauthorized, 0, err.Error(), nil)

	}

	claims, ok := tokenByte.Claims.(jwt.MapClaims)
	if !ok || !tokenByte.Valid {
		return handler.ResponseHttp(f, fiber.StatusUnauthorized, 0, "your not have permission", nil)
	}
	f.Locals("user", claims)
	return f.Next()
}
