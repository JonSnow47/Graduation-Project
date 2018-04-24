package controllers

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
)

type BaseController struct {
	beego.Controller
}

// ParseToken parse JWT token in http header.
func (c *BaseController) ParseToken() (t *jwt.Token, err error) {
	authString := c.Ctx.Input.Header("Authorization")
	beego.Debug("AuthString:", authString)

	kv := strings.Split(authString, " ")
	if len(kv) != 2 || kv[0] != "Bearer" {
		beego.Error("AuthString invalid:", authString)
		return nil, err
	}
	tokenString := kv[1]

	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("mykey"), nil
	})
	if err != nil {
		beego.Error("Parse token:", err)
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				// That‘s not even a token
				return nil, err
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// Token is either expired or not active yet
				return nil, err
			} else {
				// Couldn‘t handle this token
				return nil, err
			}
		} else {
			// Couldn‘t handle this token
			return nil, err
		}
	}
	if !token.Valid {
		beego.Error("Token invalid:", tokenString)
		return nil, err
	}
	beego.Debug("Token:", token)

	return token, nil
}
