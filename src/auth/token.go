package auth

import (
	"auth/src/models"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("SECRET_KEY"))

type Claims struct {
	UserID   uint64 `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenToken(user models.User) (string, error) {
	claims := Claims{
		UserID:   user.ID,
		Username: user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 48)),
			Issuer:    "auth_service",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, jwt.ErrTokenSignatureInvalid
			}
			return secretKey, nil
		},
	)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
