package main

import (
	_ "checkpoint-go/initial"
	_ "checkpoint-go/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.WebConfig.StaticDir["/static"] = "static"
	beego.Run()
}
