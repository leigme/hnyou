package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

//get方法登录
func (this *MainController) Get() {
	this.Data["IsIndex"] = true
	this.TplNames = "home.html"
}
