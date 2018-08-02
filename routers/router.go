package routers

import (
	"WebGo/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    beego.Get("/test", func(context *context.Context) {
		context.WriteString("hello world")
	})
}

