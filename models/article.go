package models

import (
	"github.com/jinxing-go/mysql"
)

// 定义文章模型
type Article struct {
	Id         int64      `db:"id" json:"id"`
	UserId     int64      `db:"user_id" json:"user_id"`
	Title      string     `db:"title" form:"title" json:"title"`
	Content    string     `db:"content" form:"content" json:"content"`
	Img        string     `db:"img" form:"img" json:"img"`
	Type       int        `db:"type" json:"type"`
	SeeNum     int        `db:"see_num" json:"see_num"`
	CommentNum int        `db:"comment_num" json:"comment_num"`
	Recommend  int        `db:"recommend" json:"recommend"`
	Status     int        `db:"status" json:"status"`
	CreatedAt  mysql.Time `db:"created_at" json:"created_at"`
	UpdatedAt  mysql.Time `db:"updated_at" json:"updated_at"`
}

func (*Article) TableName() string {
	return "article"
}

func (*Article) PK() string {
	return "id"
}

func (*Article) TimestampsValue() interface{} {
	return mysql.Now()
}
