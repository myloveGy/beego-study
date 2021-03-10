package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"

	"project/help"
)

type Admin struct {
	UserId        int64     `orm:"column(user_id);auto;pk" json:"user_id" form:"user_id"`
	Username      string    `orm:"column(menu_name);" form:"menu_name" json:"menu_name"`
	Password      string    `orm:"column(icons);" form:"icons" json:"icons"`
	Email         string    `orm:"column(email);" form:"email" json:"email"`
	Status        int       `orm:"column(status)" json:"status" form:"status"`
	LastLoginTime time.Time `orm:"column(last_login_time)" json:"last_login_time"`
	LastLoginIp   string    `orm:"column(last_login_ip)" json:"last_login_ip"`
	CreatedAt     time.Time `orm:"column(created_at);auto_now_add;type(datetime)" json:"created_at"`
	UpdatedAt     time.Time `orm:"column(updated_at);auto_now;type(datetime)" json:"updated_at"`
}

func (*Admin) TableName() string {
	return "admin"
}

func (a *Admin) Login(username, password, ip string) error {
	o := orm.NewOrm()

	// 查询用户
	if err := o.QueryTable(a).Filter("username", username).One(a); err != nil {
		return errors.New("用户名或者密码错误")
	}

	// 验证密码
	if !help.ValidatePassword(password, a.Password) {
		return errors.New("用户名或者密码错误")
	}

	// 验证状态
	if a.Status != 1 {
		return errors.New("对不起！你被管理员封了 ):")
	}

	// 修改登录IP
	a.LastLoginTime = time.Now()
	a.LastLoginIp = ip
	if _, err := o.Update(a, "last_login_ip", "last_login_time"); err != nil {
		return errors.New("修改登录信息失败")
	}

	return nil
}
