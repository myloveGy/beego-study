package models

// 定义分类模型
type Category struct {
	Id         int64  `orm:"column(id);auto;pk"json:"id"form:"id"`
	Pid        int64  `orm:"column(pid)"json:"pid"form:"pid"`
	Path       string `orm:"column(path)"json:"path"`
	CateName   string `orm:"column(cate_name)"json:"cate_name"form:"cate_name"`
	Sort       string `orm:"column(sort)"json:"sort"form:"sort"`
	Recommend  int    `orm:"column(recommend)"json:"recommend"form:"recommend"`
	Status     int    `orm:"column(status)"json:"status"form:"status"`
	CreateTime int64  `orm:"column(create_time)"json:"create_time"`
	CreateId   int64  `orm:"column(create_id)"json:"create_id"`
	UpdateTime int64  `orm:"column(update_time)"json:"update_time"`
	UpdateId   int64  `orm:"column(update_id)"json:"update_id"`
}

// 返回表名字
func (u *Category) TableName() string {
	return "my_category"
}