package models

import (
	"checkpoint-go/utils"
	"github.com/astaxie/beego/orm"
)

type Users struct {
	Id int
	Username string
	Password string
	Salt string
	Nickname string
	CompanyName string
	CompanyPhone string
	CompanyAddress string
	Address string
	Phone string
	Avatar string
}

func init() {
	orm.RegisterModel(new(Users))
}
// 添加用户
func AddUser(u *Users) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	user := new(Users)
	user.Username = u.Username
	user.Salt = utils.RandString(10)
	user.Password = utils.Md5(u.Password + user.Salt)
	user.Nickname = u.Nickname
	user.CompanyAddress = u.CompanyAddress
	user.CompanyName = u.CompanyName
	user.CompanyPhone = u.CompanyPhone
	user.Address = u.Address
	user.Phone = u.Phone

	return o.Insert(user)
}
// 添加用户
func AddLoginUser(username string, password string) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	user := new(Users)
	user.Username = username
	user.Salt = utils.RandString(10)
	user.Password = utils.Md5(password + user.Salt)
	return o.Insert(user)
}

// 用户名查找用户
func FindUser(username string) (Users, error) {
	o := orm.NewOrm()
	o.Using("default")
	user := Users{Username: username}
	err := o.Read(&user, "username")
	return user, err
}

// 根据id查找用户
func FindUserById(id int) ([]orm.Params, error) {
	o := orm.NewOrm()
	sql := "select id,username,nickname,avatar,company_name,company_address,company_phone,address,phone,created_at from users where id=?"
	var user []orm.Params
	_, err := o.Raw(sql, id).Values(&user)
	return user, err
}

// 编辑资料：上传头像，修改昵称
func AlterUserInfo(id int, avatar string, nickname string) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	user := Users{Id: id}
	err := o.Read(&user)
	var num int64
	if err == nil {
		user.Nickname = nickname
		user.Avatar = avatar
		num, err = o.Update(&user, "nickname", "avatar")
	}
	return num, err
}