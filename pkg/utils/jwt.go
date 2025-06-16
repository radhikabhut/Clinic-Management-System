package utils

import (
	"clinic-management/config"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(email, role string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"role":  role,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.RuntimeConfig.JWTSecret))
}

func ValidateToken(tokenStr string) (*jwt.Token, bool) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.RuntimeConfig.JWTSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, false
	}
	return token, true
}
