package routers

import (
	"beegoAdmin/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/*", &controllers.MainController{})
}
