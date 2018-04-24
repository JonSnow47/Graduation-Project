package util

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const TokenKey = "Graduation-Project"

const (
	ErrAbsent  = "token absent"  // 令牌不存在
	ErrInvalid = "token invalid" // 令牌无效
	ErrExpired = "token expired" // 令牌过期
	ErrOther   = "other error"   // 其他错误
)

func NewToken(id string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["id"] = id
	if id == "5ad996a35111cd06bd0e1e74" {
		claims["admin"] = true
	} else {
		claims["admin"] = false
	}
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token.Claims = claims

	return token.SignedString([]byte(TokenKey))
}

func ValidateToken(tokenString string) (bool, error) {
	if tokenString == "" {
		return false, errors.New(ErrAbsent)
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(TokenKey), nil
	})
	if token == nil {
		return false, errors.New(ErrInvalid)
	}
	if token.Valid {

		return true, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return false, errors.New(ErrInvalid)
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return false, errors.New(ErrExpired)
		} else {
			return false, errors.New(ErrOther)
		}
	} else {
		return false, errors.New(ErrOther)
	}
}
