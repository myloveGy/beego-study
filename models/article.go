package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// 定义文章模型
type Article struct {
	Id         int64     `orm:"column(id);auto;pk" json:"id"`
	UserId     int64     `orm:"column(user_id)" json:"user_id"`
	Title      string    `orm:"column(title);" form:"title" json:"title"`
	Content    string    `orm:"column(content);" form:"content" json:"content"`
	Img        string    `orm:"column(img);" form:"img" json:"img"`
	Type       int       `orm:"column(type);default(1)" json:"type"`
	SeeNum     int       `orm:"column(see_num)" json:"see_num"`
	CommentNum int       `orm:"column(comment_num)" json:"comment_num"`
	Recommend  int       `orm:"column(recommend)" json:"recommend"`
	Status     int       `orm:"column(status);default(1)" json:"status"`
	CreatedAt  time.Time `orm:"column(created_at);auto_now_add;type(datetime)" json:"created_at"`
	UpdatedAt  time.Time `orm:"column(updated_at);auto_now;type(datetime)" json:"updated_at"`
}

func (u *Article) TableName() string {
	return "article"
}

/**
 * GetArticle 根据条件获取文章信息
 * @param  where map[string]interface{} 查询条件
 * @param  limit int                    查询条数
 * @param  order string                 排序条件
 * @return 返回查询数据和错误信息
 */
func GetArticle(where map[string]interface{}, limit int, order string) (articles []*Article, err error) {
	qs := orm.NewOrm().QueryTable(new(Article))
	for k, v := range where {
		qs = qs.Filter(k, v)
	}
	_, err = qs.OrderBy(order).Limit(limit).All(&articles)
	return
}
