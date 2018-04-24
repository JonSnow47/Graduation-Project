package main

import (
	"log"

	"github.com/JonSnow47/Graduation-Project/blog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {
	StartServer()
}

func StartServer() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	log.SetFlags(log.Ldate | log.Lmicroseconds)
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
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
