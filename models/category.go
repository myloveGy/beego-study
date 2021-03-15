package models

import (
	"github.com/jinxing-go/mysql"

	"project/connection"
)

type Category struct {
	Id        int64      `db:"id" json:"id" form:"id"`
	Pid       int64      `db:"pid" json:"pid" form:"pid"`
	Path      string     `db:"path" json:"path"`
	CateName  string     `db:"cate_name" json:"cate_name" form:"cate_name"`
	Sort      string     `db:"sort" json:"sort" form:"sort"`
	Recommend int        `db:"recommend" json:"recommend" form:"recommend"`
	Status    int        `db:"status" json:"status" form:"status"`
	CreatedAt mysql.Time `db:"created_at" json:"created_at"`
	UpdatedAt mysql.Time `db:"updated_at" json:"updated_at"`
}

func (*Category) TableName() string {
	return "category"
}

func (*Category) PK() string {
	return "id"
}

func (*Category) TimestampsValue() interface{} {
	return mysql.Now()
}

/**
 * GetArticle 根据条件获取文章信息
 * @param  where map[string]interface{} 查询条件
 * @param  limit int                    查询条数
 * @param  order string                 排序条件
 * @return 返回查询数据和错误信息
 */
func GetArticle(where map[string]interface{}, limit int, order, by string) (articles []*Article, err error) {
	qs := connection.DB.Builder(articles)
	for k, v := range where {
		qs = qs.Where(k, v)
	}

	err = qs.OrderBy(order, by).Limit(limit).All()
	return
}
