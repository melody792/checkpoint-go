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
	// 获取用户列表
	beego.Router("/api/users", &controllers.GetUsersController{})
	// 添加用户
	beego.Router("/api/addUser", &controllers.AddUserController{})
	// 删除用户deleteUser
	beego.Router("/api/deleteUser", &controllers.DeleteUserInfoByIdController{})
	// 查找用户
	beego.Router("/api/getUserById", &controllers.GetUserByIdContrller{})
	// 用户详情
	beego.Router("/api/userInfo", &controllers.EditUserInfoController{})
}
