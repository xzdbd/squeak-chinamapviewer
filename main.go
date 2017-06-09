package main

import (
	"github.com/astaxie/beego"
	_ "github.com/xzdbd/squeak-chinamapviewer/routers"
)

func main() {
	beego.Run()
}
