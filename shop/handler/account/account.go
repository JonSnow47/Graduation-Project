/*
 * Revision History:
 *     Initial: 2018/04/28        Chen Yanchen
 */

package account

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/JonSnow47/Graduation-Project/shop/model/account"
)

// WechatLogin login with wechat permission.
func WechatLogin(c echo.Context) error {
	return nil
}

// PhoneLogin login with phone validate code.
func PhoneLogin(c echo.Context) error {
	return nil
}

// Register in web.
func Register(c echo.Context) error {
	var req struct {
		Name string `json:"name" validate:""`
		Pwd  string `json:"pwd" validate:""`
	}

	if err := c.Bind(&req); err != nil {
		log.Println("Request parametes error", err)
		return c.JSON(http.StatusBadRequest, "Request parametes error")
	}

	if err := c.Validate(&req); err != nil {
		log.Println("Request parametes error", err)
		return c.JSON(http.StatusBadRequest, "Request parametes error")
	}

	avatar, err := c.FormFile("avatar")
	if err != nil {
		return err
	}

	src, err := avatar.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// 存储头像
	dst, err := os.Create(avatar.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	err = account.AccountService.Register(req.Name, req.Pwd, dst.Name())
	if err != nil{
		return err
	}

	return c.JSON(http.StatusCreated,"")
}

// Login in web.
func Login(c echo.Context) error {
	return nil
}

// Logout delete session or close JWT.
func Loginout(c echo.Context) error {
	return nil
}
