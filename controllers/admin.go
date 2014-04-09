package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

//get方法登录
func (this *AdminController) Get() {
	this.Data["IsLogin"] = true
	this.TplNames = "admin/admin.html"
}

//post方法登录
func (this *AdminController) Post() {
	username := this.Input().Get("username")
	if "" == username {
		this.TplNames = "admin/login.html"
	} else {
		password := this.Input().Get("password")
		if "" == password {
			this.TplNames = "admin/login.html"
		} else {
			this.Data["IsUname"] = true
			this.TplNames = "admin/admin.html"
		}
	}
}

func (this *AdminController) Updata() {
	fmt.Println("调用Updata函数！")
	fmt.Println(this.Input().Get("editor"))
	this.TplNames = "admin/admin.html"
}
