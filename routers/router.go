package routers

import (
	"checkpoint-go/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// 注册接口
	beego.Router("/api/register", &controllers.RegisterController{})
	// 登录接口
	beego.Router("/api/login", &controllers.LoginController{})
	// 用户详情
	beego.Router("/api/userInfo", &controllers.EditUserInfoController{})
}
