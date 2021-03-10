package models

import (
	"time"
)

// 定义导航栏模型
type Menu struct {
	Id        int64     `orm:"column(id);auto;pk" json:"id" form:"id"`
	Pid       int64     `orm:"column(pid);" form:"pid" json:"pid"`
	MenuName  string    `orm:"column(menu_name);" form:"menu_name" json:"menu_name"`
	Icons     string    `orm:"column(icons);" form:"icons" json:"icons"`
	Url       string    `orm:"column(url);default(1)" form:"url" json:"url"`
	Status    int       `orm:"column(status)" json:"status" form:"status"`
	Sort      int       `orm:"column(sort)" json:"sort" form:"sort"`
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" json:"created_at"`
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" json:"updated_at"`
}

// 返回表名字
func (u *Menu) TableName() string {
	return "my_menu"
}
