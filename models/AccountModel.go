package models

import (
	"crypto/sha1"
	"fmt"
	"github.com/astaxie/beego/orm"
)

type AccountModel struct {
	AccountId int
	Username  string
	Password  string
}

func (this *AccountModel) GetUserInfo(username string) AccountModel {
	var res AccountModel
	o := orm.NewOrm()
	o.Raw("SELECT  account_id, username, password FROM account WHERE username = ?", username).QueryRow(&res)
	return res
}

func (this *AccountModel) AddUser(username, password string) AccountModel {

	user := AccountModel{0, username, password}
	pass := fmt.Sprintf("%x", sha1.Sum([]byte(password)))

	o := orm.NewOrm()
	res, err := o.Raw("INSERT INTO account (username, password) VALUES (?, ?)", username, pass).Exec()

	if err == nil {
		id, _ := res.LastInsertId()
		user.AccountId = int(id)
	}

	return user
}
