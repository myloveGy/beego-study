package models

import (
	"errors"
	"fmt"

	"github.com/jinxing-go/mysql"

	"project/connection"
	"project/utils"
)

type Admin struct {
	UserId        int64      `db:"user_id" json:"user_id" form:"user_id"`
	Username      string     `db:"username" form:"username" json:"username"`
	Password      string     `db:"password" form:"password" json:"password"`
	Email         string     `db:"email" form:"email" json:"email"`
	Status        int        `db:"status" json:"status" form:"status"`
	LastLoginTime mysql.Time `db:"last_login_time" json:"last_login_time"`
	LastLoginIp   string     `db:"last_login_ip" json:"last_login_ip"`
	CreatedAt     mysql.Time `db:"created_at" json:"created_at"`
	UpdatedAt     mysql.Time `db:"updated_at" json:"updated_at"`
}

func (*Admin) TableName() string {
	return "admin"
}

func (*Admin) PK() string {
	return "user_id"
}

func (*Admin) TimestampsValue() interface{} {
	return mysql.Now()
}

func (a *Admin) BeforeSave() error {

	if a.Password != "" {
		fmt.Println(a.Password, "password")
		if password, err := utils.GeneratePassword(a.Password); err != nil {
			return err
		} else {
			a.Password = password
		}
	}

	return nil
}

func (a *Admin) Login(username, password, ip string) error {
	a.Username = username

	// 查询用户
	if err := connection.DB.Find(a); err != nil {
		return errors.New("用户名或者密码错误")
	}

	// 验证密码
	if !utils.ValidatePassword(password, a.Password) {
		return errors.New("用户名或者密码错误")
	}

	// 验证状态
	if a.Status != 1 {
		return errors.New("对不起！你被管理员封了 ):")
	}

	// 修改登录IP
	if _, err := connection.DB.Update(&Admin{
		UserId:        a.UserId,
		LastLoginIp:   ip,
		LastLoginTime: mysql.Now(),
	}); err != nil {
		return errors.New("修改登录信息失败")
	}

	return nil
}
