package controllers

import (
	"github.com/astaxie/beego"
)

type EditorController struct {
	beego.Controller
}

func (this *EditorController) Get() {
	this.TplNames = "admin/editor.html"
}
