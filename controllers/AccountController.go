package controllers

import (
	"WebGo/models"
	"crypto/sha1"
	"fmt"
	"github.com/astaxie/beego"
)

type AccountController struct {
	beego.Controller
}

func (this *AccountController) Login() {
	username := this.GetString("username")
	password := this.GetString("password")

	accountModel := models.AccountModel{}
	response := make(map[string]bool)

	if info := accountModel.GetUserInfo(username); info.AccountId != 0 {
		t := fmt.Sprintf("%x", sha1.Sum([]byte(password)))
		if t == info.Password {
			this.SetSession("account_id", info.AccountId)
			this.SetSession("username", info.Username)

			response["status"] = true
			this.Data["json"] = response
			this.ServeJSON()
			return
		}
	}

	response["status"] = false
	this.Data["json"] = response
	this.ServeJSON()

}

func (this *AccountController) Register() {
	username := this.GetString("username")
	password := this.GetString("password")

	accountModel := models.AccountModel{}

	if user := accountModel.AddUser(username, password); user.AccountId != 0 {
		fmt.Println("ok")
	} else {
		fmt.Println("false")
	}

	this.Ctx.WriteString("....")
}

func (this *AccountController) Sign() {
	this.TplName = "sign.html"
}
