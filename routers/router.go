package routers

import (
	"github.com/astaxie/beego"
	"github.com/xzdbd/squeak-chinamapviewer/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
