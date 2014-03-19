package main

import (
	"hnyou/controllers"
	"github.com/astaxie/beego"
)

func main() {
	//注册路由
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})

	//启动服务
	beego.Run()
}

