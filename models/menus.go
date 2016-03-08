package models

import (
	"github.com/astaxie/beego/orm"
)

type Menus struct {
	Id         int    `orm:"column(id);auto;pk"`
	Pid        int    `orm:"column(pid);"`
	MenuName   string `orm:"column(menu_name);"`
	Icons      string `orm:"column(icons);"`
	Url        string `orm:"column(url)"`
	Status     int    `orm:"column(status)"`
	Sort       int    `orm:"column(sort)"`
	CreateTime int64  `orm:"column(create_time)"`
	CreateId   int    `orm:"column(create_id)"`
	UpdateTime int64  `orm:"column(update_time)"`
	UpdateId   int    `orm:"column(update_id)"`
}

// 初始化注册
func init() {
	orm.RegisterModel(new(Menus))
}

// 自定义表名字
func (menus *Menus) TableName() string {
	return "my_blog_menus"
}

// 获取全部数据
func MenusGetAll() (total int64, data interface{}, err error) {
	var menus []*Menus
	o := orm.NewOrm()
	total, err = o.QueryTable(new(Menus)).All(&menus)
	if err == nil {
		data = menus
	}

	return
}
