package main

import (
	_ "rbac-beego-api/routers"

	beego "github.com/beego/beego/v2/server/web"
	"rbac-beego-api/database"
)

func main() {
	database.InitDatabase()
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
