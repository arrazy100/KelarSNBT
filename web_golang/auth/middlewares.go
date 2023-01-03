package auth

import (
	"errors"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Data map[string]interface{} `json:"data"`
	jwt.StandardClaims
}

func AuthAdminMiddleware(ctx *fiber.Ctx) error {
	publicKey, err := os.ReadFile(os.Getenv("PUBLIC_KEY"))

	if err != nil {
		return err
	}

	tokenString := string(ctx.Request().Header.Peek("Authorization"))

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwt.ParseRSAPublicKeyFromPEM(publicKey)
	})

	if err != nil {
		return err
	}

	haveAccess := false

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		roles := claims.Data["roles"].([]interface{})

		for _, v := range roles {
			if v == "admin" {
				haveAccess = true
				break
			}

			haveAccess = false
		}

		if !haveAccess {
			return errors.New("access not granted")
		}
	} else {
		ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "Invalid token"})
		return errors.New("invalid token")
	}

	return ctx.Next()
}
