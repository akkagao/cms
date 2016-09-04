package main

import (
	"github.com/astaxie/beego"

	_ "cms/src/routers"
	_ "cms/src/service"
)

func main() {
	beego.SetLevel(beego.LevelDebug)
	beego.SetLogger("console", "")
	beego.Run()
}
