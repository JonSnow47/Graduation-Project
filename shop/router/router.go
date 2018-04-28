/*
 * Revision History:
 *     Initial: 2018/04/28        Chen Yanchen
 */

package router

import (
	"github.com/JonSnow47/Graduation-Project/shop/handler"
	"github.com/labstack/echo"
)

func Init(e *echo.Echo) *echo.Echo {
	e.GET("/", handler.Home)
	e.POST("/login", handler.Login)
	e.POST("/save", handler.Save)
	return e
}
