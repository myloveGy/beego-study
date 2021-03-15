package models

import (
	"github.com/jinxing-go/mysql"
)

// 定义分类模型
type Image struct {
	ImageId     int64      `db:"image_id" json:"image_id" form:"image_id"`
	UserId      int64      `db:"user_id" json:"user_id" form:"user_id"`
	Title       string     `db:"title" json:"title" form:"title"`
	Description string     `db:"description" json:"description" form:"description"`
	Url         string     `db:"url" json:"url" form:"url"`
	Type        int        `db:"type" json:"type" form:"type"`
	Sort        int32      `db:"sort" json:"sort" form:"sort"`
	Status      int        `db:"status" json:"status" form:"status"`
	CreatedAt   mysql.Time `db:"created_at" json:"created_at"`
	UpdatedAt   mysql.Time `db:"updated_at" json:"updated_at"`
}

// 返回表名字
func (*Image) TableName() string {
	return "image"
}

func (*Image) PK() string {
	return "image_id"
}

func (*Image) TimestampsValue() interface{} {
	return mysql.Now()
}
