package models

import (
	"github.com/astaxie/beego/orm"
)

// 定义导航栏模型
type Menu struct {
	Id         int64  `orm:"column(id);auto;pk"json:"id"`
	Pid        int64  `orm:"column(pid);"form:"title"json:"pid"`
	MenuName   string `orm:"column(menu_name);"form:"content"json:"menu_name"`
	Icons      string `orm:"column(icons);"form:"img"json:"icons"`
	Url        string `orm:"column(url);default(1)"json:"url"`
	Status     int    `orm:"column(status)"json:"status"`
	CreateTime int64  `orm:"column(create_time)"json:"create_time"`
	CreateId   int64  `orm:"column(create_id)"json:"create_id"`
	UpdateTime int64  `orm:"column(update_time)"json:"update_time"`
	UpdateId   int64  `orm:"column(update_id)"json:"update_id"`
}

func (u *Menu) TableName() string {
	return "my_menu"
}

// 初始化注册
func init() {
	orm.RegisterModel(new(Menu))
}

// 查询数据
func SearchMenu(query map[string]interface{}, limit int64, offset int64, order string) (array []*Menu, err error) {
	qs := orm.NewOrm().QueryTable(new(Menu))
	for k, v := range query {
		qs = qs.Filter(k, v)
	}

	qs = qs.OrderBy(order).Limit(limit, offset)
	_, err = qs.All(&array)
	return
}
