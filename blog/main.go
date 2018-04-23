package main

import (
	_ "github.com/JonSnow47/Graduation-Project/blog/routers"

	"github.com/astaxie/beego"
	"log"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {
	StartServer()
}

func StartServer()  {
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
		AllowCredentials: false,
	}))
	beego.InsertFilter("/*", beego.BeforeRouter, filters.LoginFilter)
	beego.Run()
}