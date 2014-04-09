package main

import (
	"github.com/astaxie/beego"
	"hnyou/controllers"
)

func main() {
	//注册路由
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/admin", &controllers.AdminController{})
	beego.Router("/updata", &controllers.AdminController{}, "*:Updata")
	beego.Router("/editor", &controllers.EditorController{})

	//启动服务
	beego.Run()
}
