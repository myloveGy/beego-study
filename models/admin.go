package models

import (
	// "app/controllers"
	"crypto/sha1"
	"fmt"
	"github.com/astaxie/beego/orm"
	"io"
	"time"
)

type Admin struct {
	Id          int64  `orm:"column(id);auto;pk"`
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

// 获取全部数据
func AdminGetAll(query map[string]interface{}, offset int64, limit int64, order string) (total int64, num int, data interface{}, err error) {
	var admin []*Admin
	o := orm.NewOrm()
	qu := o.QueryTable(new(Admin))

	for k, v := range query {
		qu = qu.Filter(k, v)
	}

	total, _ = qu.Count()
	_, err = qu.OrderBy(order).Limit(limit, offset).All(&admin)
	if err == nil {
		num = len(admin)
		data = admin
	}
	return
}

/**
 * [AdminLogout description] 根据ID修改管理员最后登录信息
 * @param {[type]} id   int    [description] 唯一ID
 * @param {[type]} ip   string [description] IP
 * @param {[type]} time int64  [description] 最后登录的时间
 */
func AdminLogout(id int64, ip string, time int64) (err error) {
	o := orm.NewOrm()
	admin := Admin{Id: id}
	if err = o.Read(&admin); err == nil {
		admin.LastIp = ip
		admin.LastTime = time
		_, err = o.Update(&admin)
	}

	return
}

// 新增数据
func AdminInsert(admin *Admin) (InsertId int64, err error) {
	o := orm.NewOrm()
	time := time.Now().Unix()
	admin.CreateTime = time
	admin.LastTime = time

	// 密码加密
	t := sha1.New()
	io.WriteString(t, admin.Password)
	admin.Password = fmt.Sprintf("%x", t.Sum(nil))
	InsertId, err = o.Insert(admin)
	return
}

// 修改数据
func AdminSave(admin *Admin) (isHave bool, RowCount int64, err error) {
	o := orm.NewOrm()
	update := Admin{Id: admin.Id}
	isHave = false
	if o.Read(&update) == nil {
		isHave = true
		// 密码加密
		if admin.Password != "" {
			t := sha1.New()
			io.WriteString(t, admin.Password)
			admin.Password = fmt.Sprintf("%x", t.Sum(nil))
		} else {
			admin.Password = update.Password
		}

		RowCount, err = o.Update(admin)
	}

	return
}

// 删除数据
func AdminDelete(Id int64) (isHave bool, RowCount int64, err error) {
	o := orm.NewOrm()
	admin := Admin{Id: Id}
	err = o.Read(&admin)
	isHave = false
	if err == nil {
		isHave = true
		RowCount, err = o.Delete(&admin)
	}

	return
}
