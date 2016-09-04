package routers

import (
	"cms/src/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Info("init routers start ...")

	beego.Router("/", &controllers.MainController{}, "*:Index")
	beego.Router("/welcome", &controllers.MainController{}, "*:Welcome")
	beego.Router("/leftMenu", &controllers.MainController{}, "*:LeftMenu")
	beego.Router("/header", &controllers.MainController{}, "*:Header")
	beego.Router("/loadMenu", &controllers.MainController{}, "*:LoadMenu")
	// beego.Router("/login", &controllers.MainController{}, "*:Login")
	// beego.Router("/tologin", &controllers.MainController{}, "*:Tologin")
	// beego.Router("/loginpage", &controllers.MainController{}, "*:Loginpage")
	//自动绑定映射关系
	beego.AutoRouter(&controllers.AdmUserController{})
	beego.AutoRouter(&controllers.LoginController{})
	beego.AutoRouter(&controllers.RoleController{})
	beego.AutoRouter(&controllers.AdmUserGroupController{})

	beego.Info("init routers end.")
}
