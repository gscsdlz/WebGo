package routers

import (
	"WebGo/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/sign", &controllers.AccountController{}, "get:Sign")
	beego.Router("/login", &controllers.AccountController{}, "post:Login")
	beego.Router("/register", &controllers.AccountController{}, "post:Register")

	beego.Router("/chat", &controllers.MainController{}, "get:Chat")
	//beego.Router("/chat", &controllers.MainController{}, "post:Chat")

	beego.Router("/", &controllers.MainController{}, "get:Index")
}
