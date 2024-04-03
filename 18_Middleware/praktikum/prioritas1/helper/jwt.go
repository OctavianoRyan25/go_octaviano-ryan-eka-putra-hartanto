package helper

import (
	"errors"
	"prioritas1gorm/constants"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

var secretKey = constants.SECRET_JWT

func GenerateToken(id uint, email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["email"] = email

	// Membuat token JWT dengan klaim yang diberikan
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// VerifyToken memverifikasi token JWT dari konteks Echo
func VerifyToken(c echo.Context) (jwt.MapClaims, error) {
	errResponse := errors.New("authentication failed")
	headerToken := c.Request().Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer ")

	if !bearer {
		return nil, errResponse
	}

	stringToken := strings.Split(headerToken, " ")[1]
	token, err := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errResponse
}
