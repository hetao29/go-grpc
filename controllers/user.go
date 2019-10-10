package controllers

import (
	"github.com/astaxie/beego"
	"user"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	userinfo := user.NewUserInfo()
	userinfo.Set("hahah")
	v := this.GetSession("asta")
	if v == nil {
		this.SetSession("asta", int(1))
	} else {
		this.SetSession("asta", v.(int)+1)
	}

	this.Data["Ct"] = this.GetSession("asta");
	this.Data["name"] = this.GetString("name")
}
