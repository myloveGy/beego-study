package models

import (
	"github.com/astaxie/beego/orm"
	// "strings"
	"time"
)

type Menus struct {
	Id         int64  `orm:"column(id);auto;pk"`
	Pid        int64  `orm:"column(pid);"`
	MenuName   string `orm:"column(menu_name);"`
	Icons      string `orm:"column(icons);"`
	Url        string `orm:"column(url)"`
	Status     int    `orm:"column(status)"`
	Sort       int    `orm:"column(sort)"`
	CreateTime int64  `orm:"column(create_time)"`
	CreateId   int64  `orm:"column(create_id)"`
	UpdateTime int64  `orm:"column(update_time)"`
	UpdateId   int64  `orm:"column(update_id)"`
}

// 初始化注册
func init() {
	orm.RegisterModel(new(Menus))
}

// 自定义表名字
func (menus *Menus) TableName() string {
	return "my_blog_menus"
}

// 查询全部
func GetAllMenus() (menus []*Menus, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable(new(Menus)).Filter("status", 1).OrderBy("id").All(&menus)
	return
}

// 获取全部数据
func MenusGetAll(query map[string]interface{}, offset int64, limit int64, order string) (total int64, num int, data interface{}, err error) {
	var menus []*Menus
	o := orm.NewOrm()
	qu := o.QueryTable(new(Menus))

	for k, v := range query {
		qu = qu.Filter(k, v)
	}

	total, _ = qu.Count()
	_, err = qu.OrderBy(order).Limit(limit, offset).All(&menus)
	if err == nil {
		num = len(menus)
		data = menus
	}
	return
}

// 新增数据
func MenusInsert(menus *Menus) (InsertId int64, err error) {
	o := orm.NewOrm()
	time := time.Now().Unix()
	menus.CreateTime = time
	menus.UpdateTime = time
	menus.CreateId = menus.UpdateId
	InsertId, err = o.Insert(menus)
	return
}

// 修改数据
func MenusSave(menus *Menus) (isHave bool, RowCount int64, err error) {
	o := orm.NewOrm()
	update := Menus{Id: menus.Id}
	isHave = false
	if o.Read(&update) == nil {
		isHave = true
		menus.UpdateTime = time.Now().Unix()
		RowCount, err = o.Update(menus)
	}

	return
}

// 删除数据
func MenusDelete(Id int64) (isHave bool, RowCount int64, err error) {
	o := orm.NewOrm()
	menus := Menus{Id: Id}
	err = o.Read(&menus)
	isHave = false
	if err == nil {
		isHave = true
		RowCount, err = o.Delete(&menus)
	}

	return
}
