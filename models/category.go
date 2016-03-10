package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Category struct {
	Id         int64  `orm:"column(id);auto;pk"`
	Pid        int64  `orm:"column(pid);"`
	Path       string `orm:"column(path);"`
	CateName   string `orm:"column(cate_name);"`
	Sort       int    `orm:"column(sort)"`
	Recommend  int    `orm:"column(recommend)"`
	Status     int    `orm:"column(status)"`
	CreateTime int64  `orm:"column(create_time)"`
	CreateId   int64  `orm:"column(create_id)"`
	UpdateTime int64  `orm:"column(update_time)"`
	UpdateId   int64  `orm:"column(update_id)"`
}

// 初始化注册
func init() {
	orm.RegisterModel(new(Category))
}

// 自定义表名字
func (cate *Category) TableName() string {
	return "my_blog_category"
}

// 获取全部数据
func CategoryGetAll(query map[string]interface{}, offset int64, limit int64, order string) (total int64, num int, data interface{}, err error) {
	var cate []*Category
	o := orm.NewOrm()
	qu := o.QueryTable(new(Category))

	for k, v := range query {
		qu = qu.Filter(k, v)
	}

	total, _ = qu.Count()
	_, err = qu.OrderBy(order).Limit(limit, offset).All(&cate)
	if err == nil {
		num = len(cate)
		data = cate
	}
	return
}

// 新增数据
func CategoryInsert(cate *Category) (InsertId int64, err error) {
	o := orm.NewOrm()
	time := time.Now().Unix()
	cate.CreateTime = time
	cate.UpdateTime = time
	cate.CreateId = cate.UpdateId
	InsertId, err = o.Insert(cate)
	return
}

// 修改数据
func CategorySave(cate *Category) (isHave bool, RowCount int64, err error) {
	o := orm.NewOrm()
	update := Category{Id: cate.Id}
	isHave = false
	if o.Read(&update) == nil {
		isHave = true
		RowCount, err = o.Update(cate)
	}

	return
}

// 删除数据
func CategoryDelete(Id int64) (isHave bool, RowCount int64, err error) {
	o := orm.NewOrm()
	cate := Category{Id: Id}
	err = o.Read(&cate)
	isHave = false
	if err == nil {
		isHave = true
		RowCount, err = o.Delete(&cate)
	}

	return
}
