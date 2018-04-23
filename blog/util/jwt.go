package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const tokenKey = "Graduation-Project"

func NewToken(id string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["admin"] = id
	if id == "5ad996a35111cd06bd0e1e74" {
		claims["admin"] = true
	} else {
		claims["admin"] = false
	}
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token.Claims = claims

	return token.SignedString([]byte(tokenKey))
}
