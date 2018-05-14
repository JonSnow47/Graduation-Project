/*
 * Revision History:
 *     Initial: 2018/05/02        Chen Yanchen
 */

package util

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const jwtKey = "sda-wefsdvz=weuhf;awoe[ga/.sdakfnzcv"

func JwtToken(name string, admin bool, level int8) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = name
	claims["admin"] = admin
	claims["level"] = level
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}
	return t, err
}
