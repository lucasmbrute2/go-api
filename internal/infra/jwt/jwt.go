package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	Secret string
}

type jwtCustomClaims struct {
	UserId  int `json:"userId"`
	jwt.RegisteredClaims
}

func (j *JWT) Generate(userId int) (string, error) {
	claims := &jwtCustomClaims{
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(j.Secret))

	if err != nil {
		return "", err
	}

	return t, nil
}