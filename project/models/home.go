package models

import "github.com/astaxie/beego/orm"

// 定义文章模型
type Article struct {
	Id         int64  `orm:"column(id);auto;pk"`
	Title      string `orm:"column(title);"`
	Content    string `orm:"column(Content);"`
	Img        string `orm:"column(img);"`
	Type       string `orm:"column(type)"`
	SeeNum     string `orm:"column(see_num)"`
	CommentNum string `orm:"column(comment_num)"`
	Recommend  int    `orm:"column(recommend)"`
	Status     int    `orm:"column(status)"`
	CreateTime int64  `orm:"column(create_time)"`
	CreateId   int64  `orm:"column(create_id)"`
	UpdateTime string `orm:"column(update_time)"`
	UpdateId   int64  `orm:"column(update_id)"`
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
