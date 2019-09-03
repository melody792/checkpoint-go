package controllers

import (
	. "checkpoint-go/models"
	"checkpoint-go/utils"
	"fmt"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

// 注册
type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")
	if !utils.CheckUsername(username) {
		c.Data["json"] = map[string]interface{}{"code": 0, "msg": "用户名格式错误"}
		c.ServeJSON()
		return
	}
	id, err := AddLoginUser(username, password)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 0, "msg": "注册失败", "data": err}
	} else {
		c.Data["json"] = map[string]interface{}{"code": 1, "msg": fmt.Sprintf("[%d] ", id) + "注册成功"}
	}
	c.ServeJSON()
}

// 登录
type LoginController struct {
	beego.Controller
}

func (c *LoginController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")
	if username == "" || password == "" {
		c.Data["json"] = map[string]interface{}{"code": 0, "msg": "invalid request"}
	}
	user, err := FindUser(username)
	if err != nil {
		fmt.Println(username, password)
		c.Data["json"] = map[string]interface{}{"code": 0, "msg": "用户不存在"}
	} else {
		pwd := utils.Md5(password + user.Salt)
		if pwd == user.Password {
			c.SetSession("username", username)
			c.Data["json"] = map[string]interface{}{"code": 1, "msg": "用户[" + user.Username + "]登录成功"}
		} else {
			c.Data["json"] = map[string]interface{}{"code": 0, "msg": "密码错误"}
		}
	}
	c.ServeJSON()
}

// 修改资料：上传头像，修改昵称 & 获取用户资料
type EditUserInfoController struct {
	beego.Controller
}

func (c *EditUserInfoController) Get() {
	user := c.GetSession("username")
	if user == nil {
		c.Data["json"] = map[string]interface{}{"code": 1001, "msg": "未登录"}
		c.ServeJSON()
		return
	}
	id, err := c.GetInt("id")
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 0, "msg": "缺少id"}
		c.ServeJSON()
		return
	}
	userInfo, err := FindUserById(id)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"code": 0,
			"msg":  "请求失败",
			"data": err,
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"code": 1,
			"msg":  "请求成功",
			"data": userInfo,
		}
	}
	c.ServeJSON()
}

func (c *EditUserInfoController) Post() {
	user := c.GetSession("username")
	if user == nil {
		c.Data["json"] = map[string]interface{}{"code": 1001, "msg": "未登录"}
		c.ServeJSON()
		return
	}
	id, err := c.GetInt("id")
	avatar := c.GetString("avatar")
	nickname := c.GetString("nickname")
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 0, "msg": "缺少id"}
		c.ServeJSON()
		return
	}
	num, err := AlterUserInfo(id, avatar, nickname)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"code": 0,
			"msg":  "请求失败",
			"data": err,
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"code": 1,
			"msg":  "请求成功",
			"data": num,
		}
	}
	c.ServeJSON()
}
