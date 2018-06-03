package main

import (
	"fmt"
	"log"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"

	"github.com/JonSnow47/Graduation-Project/blog/routers"
)

func main() {
	startServer()
}

func startServer() {
	// set the log format.
	log.SetFlags(log.Ldate | log.Lmicroseconds)

	beego.BConfig.WebConfig.Session.SessionOn = true
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/static"] = "static"
	}

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	beego.InsertFilter("/*", beego.BeforeRouter, routers.LoginFilter)
	beego.Run()
}

func init() {
	var first, last string
	fmt.Println("Who am I?")
	fmt.Scanf("%s %s", &first, &last)
	if first != "Jon" || last != "Snow" {
		os.Exit(0)
	}
}
