/*
 * Revision History:
 *     Initial: 2018/04/26        Chen Yanchen
 */

package handler

import (
	"github.com/labstack/echo"
	"io"
	"log"
	"net/http"
	"os"
)

func Login(c echo.Context) error {
	var req struct {
		Name string `json:"name"`
		Pwd  string `json:"pwd"`
	}

	if err := c.Bind(&req); err != nil {
		log.Println(err)
		return err
	}

	if req.Name == "Jon Snow" && req.Pwd == "123456" {
		log.Println("Login.")
		return c.JSON(http.StatusOK, "Login.")
	}
	return c.JSON(http.StatusOK, "Incorrect name or password.")
}

func Home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// e.POST("/save", save)
func Save(c echo.Context) error {
	// 获取 name 和 email 的值
	name := c.FormValue("name")
	email := c.FormValue("email")
	avatar, err := c.FormFile("avatar")
	if err != nil {
		return err
	}

	src, err := avatar.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(avatar.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, "<b>Thank you! "+name+email+"</b>")
}
