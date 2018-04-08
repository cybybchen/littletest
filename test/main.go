package main

import (
	"github.com/astaxie/beego"
	"src_test/controllers"
	_ "test/routers"
)

func main() {
	beego.Run()
	controllers.File()
}
