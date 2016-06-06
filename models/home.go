package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// 定义文章模型
type Article struct {
	Id         int64  `orm:"column(id);auto;pk"json:"id"`
	Title      string `orm:"column(title);"form:"title"json:"title"`
	Content    string `orm:"column(content);"form:"content"json:"content"`
	Img        string `orm:"column(img);"form:"img"json:"img"`
	Type       int    `orm:"column(type);default(1)"json:"type"`
	SeeNum     int    `orm:"column(see_num)"json:"see_num"`
	CommentNum int    `orm:"column(comment_num)"json:"comment_num"`
	Recommend  int    `orm:"column(recommend)"json:"recommend"`
	Status     int    `orm:"column(status);default(1)"json:"status"`
	CreateTime int64  `orm:"column(create_time)"json:"create_time"`
	CreateId   int64  `orm:"column(create_id)"json:"create_id"`
	UpdateTime int64  `orm:"column(update_time)"json:"update_time"`
	UpdateId   int64  `orm:"column(update_id)"json:"update_id"`
}

func (u *Article) TableName() string {
	return "my_article"
}

// 初始化注册
func init() {
	orm.RegisterModel(new(Article))
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

/**
 * AddArticle 新增文章信息
 * @param  Article article 文章对象
 * @return Id 和 错误信息
 */
func AddArticle(article *Article) (Id int64, err error) {
	o := orm.NewOrm()
	// 默认值
	article.CreateTime = time.Now().Unix()
	article.UpdateTime = article.CreateTime
	article.Type = 1
	article.Status = 1
	Id, err = o.Insert(article)
	return
}
