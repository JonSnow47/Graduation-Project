/*
 * Revision History:
 *     Initial: 2018/04/26        Chen Yanchen
 */

package conf

import (
	"github.com/labstack/echo"
)

func Config()  {
	e := echo.New()

	// 静态资源配置
	e.Static("/static","static")
}