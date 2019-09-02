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
	// 添加&获取文档目录接口
	// beego.Router("/api/docContent", &controllers.DocContentController{})
	// 用户资料编辑接口
	beego.Router("/api/userInfo", &controllers.EditUserInfoController{})
}
