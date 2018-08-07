package main

import (
	_ "WebGo/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "root:1234@tcp(127.0.0.1:3306)/we_chat?charset=utf8", 30)

	/*	var FilterUser = func(ctx *context.Context) {

			allowURI := []string {
				"/login", "/sign", "/register",
			}

			_, ok := ctx.Input.Session("account_id").(int)
			if !ok  {
				allowSig := false
				for _, uri := range allowURI  {
					if ctx.Request.RequestURI == uri {
						allowSig = true
						break
					}
				}
				if !allowSig {
					ctx.Redirect(302, "/sign")
				}
			}
		}

		beego.InsertFilter("/*",beego.BeforeRouter,FilterUser)*/
}

func main() {

	beego.Run()
}
