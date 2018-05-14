/*
 * Revision History:
 *     Initial: 2018/04/28        Chen Yanchen
 */

package router

import (
	"github.com/labstack/echo"

	"github.com/JonSnow47/Graduation-Project/shop/handler"
	"github.com/JonSnow47/Graduation-Project/shop/handler/account"
)

func Init(e *echo.Echo) *echo.Echo {
	e.GET("/", handler.Home)
	e.POST("/login", handler.Login)
	e.POST("/save", handler.Save)

	e.POST("/account/register", account.Register)
	e.POST("/account/login", account.Login)
	return e
}
