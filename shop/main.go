/*
 * Revision History:
 *     Initial: 2018/04/28        Chen Yanchen
 */

package main

import (
	"log"

	"github.com/JonSnow47/Graduation-Project/shop/router"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func init() {

}

func main() {
	// set the log format.
	log.SetFlags(log.Ldate | log.Lmicroseconds)

	e := echo.New()

	// config
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e = router.Init(e)
	/*	s := &http.Server{
			Addr:         ":1323",
			ReadTimeout:  20 * time.Minute,
			WriteTimeout: 20 * time.Minute,
		}
		e.Logger.Fatal(e.StartServer(s))*/

	e.Logger.Fatal(e.Start(":1323"))
}
