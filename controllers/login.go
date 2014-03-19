package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

//get方法登录
func (this *LoginController) Get() {
	this.Data["IsLogin"]=true
	this.TplNames = "login.tpl"
}

//post方法登录
func (this *LoginController) Post() {
	this.Data["IsUname"] = true
	this.TplNames = "edit.tpl"
}