package routers

import (
	"blognew1/controllers"
	"github.com/astaxie/beego"
)

// 注册路由
func init() {
	// 登陆路由 beego.Router("/login",&controllers.LoginController{})
	// 分类路由 beego.Router("/category",&controllers.Category{}）
    beego.Router("/", &controllers.MainController{})
}
