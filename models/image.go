package models

import (
	"github.com/astaxie/beego/orm"
	// "strings"
	"time"
)

type Image struct {
	Id         int64  `orm:"column(id);auto;pk"`
	Title      string `orm:"column(title);"`
	Desc       string `orm:"column(desc);"`
	Url        string `orm:"column(url)"`
	Type       int    `orm:"column(type)"`
	Status     int    `orm:"column(status)"`
	Sort       int    `orm:"column(sort)"`
	CreateTime int64  `orm:"column(create_time)"`
	CreateId   int64  `orm:"column(create_id)"`
	UpdateTime int64  `orm:"column(update_time)"`
	UpdateId   int64  `orm:"column(update_id)"`
}

// 初始化注册
func init() {
	orm.RegisterModel(new(Image))
}

// 自定义表名字
func (Image *Image) TableName() string {
	return "my_blog_Image"
}

// 获取全部数据
func ImageGetAll(query map[string]interface{}, offset int64, limit int64, order string) (total int64, num int, data interface{}, err error) {
	var image []*Image
	o := orm.NewOrm()
	qu := o.QueryTable(new(Image))

	for k, v := range query {
		qu = qu.Filter(k, v)
	}

	total, _ = qu.Count()
	_, err = qu.OrderBy(order).Limit(limit, offset).All(&image)
	if err == nil {
		num = len(image)
		data = image
	}
	return
}

// 新增数据
func ImageInsert(image *Image) (InsertId int64, err error) {
	o := orm.NewOrm()
	time := time.Now().Unix()
	image.CreateTime = time
	image.UpdateTime = time
	image.CreateId = image.UpdateId
	InsertId, err = o.Insert(image)
	return
}

// 修改数据
func ImageSave(image *Image) (isHave bool, RowCount int64, err error) {
	o := orm.NewOrm()
	update := Image{Id: image.Id}
	isHave = false
	if o.Read(&update) == nil {
		isHave = true
		image.UpdateTime = time.Now().Unix()
		RowCount, err = o.Update(image)
	}

	return
}

// 删除数据
func ImageDelete(Id int64) (isHave bool, RowCount int64, err error) {
	o := orm.NewOrm()
	image := Image{Id: Id}
	err = o.Read(&image)
	isHave = false
	if err == nil {
		isHave = true
		RowCount, err = o.Delete(&image)
	}

	return
}
