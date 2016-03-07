package models

import (
	"crypto/sha1"
	"fmt"
	"github.com/astaxie/beego/orm"
	"io"
)

type Admin struct {
	Id          int    `orm:"column(id);auto;pk"`
	Username    string `orm:"column(username);"`
	Password    string `orm:"column(password);"`
	Email       string `orm:"column(email);"`
	Power       string `orm:"column(power)"`
	AutoKey     string `orm:"column(auto_key)"`
	AccessToken string `orm:"column(access_token)"`
	Status      int    `orm:"column(status)"`
	CreateTime  int64  `orm:"column(create_time)"`
	LastTime    int64  `orm:"column(last_time)"`
	LastIp      string `orm:"column(last_ip)"`
}

// 初始化注册
func init() {
	orm.RegisterModel(new(Admin))
}

// 自定义表名字
func (admin *Admin) TableName() string {
	return "my_blog_admin"
}

/**
 * AdminGetOne 查询用户是否存在
 * @param string username 用户名
 * @param string password 用户密码
 * @return bool  admin 查询到用户数据返回true 和 admin 数据
 */
func AdminGetOne(username, password string) (IsTrue bool, admin Admin) {
	t := sha1.New()
	io.WriteString(t, password)
	pass := fmt.Sprintf("%x", t.Sum(nil))
	IsTrue = false
	o := orm.NewOrm()
	err := o.QueryTable(new(Admin)).Filter("username", username).Filter("password", pass).One(&admin)
	if err == nil {
		IsTrue = true
	}

	return
}

/**
 * [AdminLogout description] 根据ID修改管理员最后登录信息
 * @param {[type]} id   int    [description] 唯一ID
 * @param {[type]} ip   string [description] IP
 * @param {[type]} time int64  [description] 最后登录的时间
 */
func AdminLogout(id int, ip string, time int64) (err error) {
	o := orm.NewOrm()
	admin := Admin{Id: id}
	if err = o.Read(&admin); err == nil {
		admin.LastIp = ip
		admin.LastTime = time
		_, err = o.Update(&admin)
	}

	return
}
