/*
 * Revision History:
 *     Initial: 2018/04/28        Chen Yanchen
 */

package util

import (
	"golang.org/x/crypto/bcrypt"
)

type pwdServiceProvider struct{}

var PwdService *pwdServiceProvider

func (*pwdServiceProvider) Generator(pwd string) (string, error) {
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashPwd), nil
}

func (*pwdServiceProvider) Comparator(hashpwd []byte, pwd string) bool {
	err := bcrypt.CompareHashAndPassword(hashpwd, []byte(pwd))
	if err == nil {
		return true
	}
	return false
}
