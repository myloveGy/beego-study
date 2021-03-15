package models

import (
	"github.com/jinxing-go/mysql"
)

// 定义导航栏模型
type Menu struct {
	Id        int64      `db:"id" json:"id" form:"id"`
	Pid       int64      `db:"pid" form:"pid" json:"pid"`
	MenuName  string     `db:"menu_name" form:"menu_name" json:"menu_name"`
	Icons     string     `db:"icons" form:"icons" json:"icons"`
	Url       string     `db:"url" form:"url" json:"url"`
	Status    int        `db:"status" json:"status" form:"status"`
	Sort      int        `db:"sort" json:"sort" form:"sort"`
	CreatedAt mysql.Time `db:"created_at" json:"created_at"`
	UpdatedAt mysql.Time `db:"updated_at" json:"updated_at"`
}

// 返回表名字
func (*Menu) TableName() string {
	return "menu"
}

func (*Menu) PK() string {
	return "id"
}

func (*Menu) TimestampsValue() interface{} {
	return mysql.Now()
}
