package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Category struct {
	Id        int64     `orm:"column(id);auto;pk" json:"id" form:"id"`
	Pid       int64     `orm:"column(pid)" json:"pid" form:"pid"`
	Path      string    `orm:"column(path)" json:"path"`
	CateName  string    `orm:"column(cate_name)" json:"cate_name" form:"cate_name"`
	Sort      string    `orm:"column(sort)" json:"sort" form:"sort"`
	Recommend int       `orm:"column(recommend)" json:"recommend" form:"recommend"`
	Status    int       `orm:"column(status)" json:"status" form:"status"`
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" json:"created_at"`
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" json:"updated_at"`
}

func (*Category) TableName() string {
	return "category"
}

func (*Category) PK() string {
	return "id"
}

// 初始化注册
func init() {
	orm.RegisterModel(new(Category))
}
