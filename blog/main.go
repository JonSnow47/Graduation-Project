package main

import (
	_ "github.com/JonSnow47/Graduation-Project/blog/routers"

	"github.com/astaxie/beego"
	"log"
)

func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds)

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
