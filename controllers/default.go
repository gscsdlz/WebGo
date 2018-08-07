package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Index() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}

func (this *MainController) Chat() {
	ws := websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}

	conn, _ := ws.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	Add(conn)
}
