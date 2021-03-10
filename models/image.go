package models

import (
	"time"
)

// 定义分类模型
type Image struct {
	ImageId     int64     `orm:"column(image_id); auto;pk" json:"image_id" form:"image_id"`
	UserId      int64     `orm:"column(user_id)" json:"user_id" form:"user_id"`
	Title       string    `orm:"column(title)" json:"title" form:"title"`
	Description string    `orm:"column(description)" json:"description" form:"description"`
	Url         string    `orm:"column(url)" json:"url" form:"url"`
	Type        int       `orm:"column(type)" json:"type" form:"type"`
	Sort        int32     `orm:"column(sort)" json:"sort" form:"sort"`
	Status      int       `orm:"column(status);default(1)" json:"status" form:"status"`
	CreatedAt   time.Time `orm:"column(created_at);auto_now_add;type(datetime)" json:"created_at"`
	UpdatedAt   time.Time `orm:"column(updated_at);auto_now;type(datetime)" json:"updated_at"`
}

// 返回表名字
func (*Image) TableName() string {
	return "image"
}
