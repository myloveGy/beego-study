package models

import (
	"errors"
	"reflect"

	"github.com/astaxie/beego/orm"
)

// 初始化注册
func init() {
	orm.RegisterModel(new(Menu))
	orm.RegisterModel(new(Category))
}

// 查询对象信息
type Query struct {
	Table   string
	Where   map[string]interface{}
	IStart  int64
	ILength int64
	Order   string
}

// 其他查询对象信息
type QueryOther struct {
	Table  string
	Where  map[string]interface{}
	Limit  int64
	Offset int64
	Order  string
}

// 获取所有数据 表 查询条件 数据条数 开始位置 排序
func FindAll(array interface{}, query Query) (total, count int64, err error) {
	// 获取查询对象
	qs := orm.NewOrm().QueryTable(query.Table)

	// 拼接查询条件
	for k, v := range query.Where {
		qs = qs.Filter(k, v)
	}

	// 查询数据总条数
	total, err = qs.Count()
	if err == nil {
		// 查询出制定数据条数
		qs = qs.OrderBy(query.Order).Limit(query.ILength, query.IStart)
		count, err = qs.All(array)
	}

	return
}

// 查询数据
func Find(query QueryOther) orm.QuerySeter {
	// 获取查询对象
	qs := orm.NewOrm().QueryTable(query.Table)

	// 判断查询条件
	if where := query.Where; where != nil {
		for k, v := range where {
			qs = qs.Filter(k, v)
		}
	}

	// 判断查询条数 和查询的起始位置
	if limit := query.Limit; limit > 0 {
		qs = qs.Limit(limit).Offset(query.Offset)
	}

	// 判断排序条件
	if order := query.Order; order != "" {
		qs = qs.OrderBy(order)
	}

	return qs
}

// 查询数据所有数据
func All(data interface{}, query QueryOther) (total int64, err error) {
	qs := Find(query)
	// 查询数据返回
	total, err = qs.All(data)
	return
}

// 查询单条数据
func One(data interface{}, query QueryOther) error {
	qs := Find(query)
	return qs.One(data)
}

// 新增数据
func Insert(object interface{}) (id int64, err error) {
	v := reflect.ValueOf(object)
	f := v.MethodByName("BeforeInsert")
	// 执行新增之前
	if f.IsValid() {
		m := f.Call([]reflect.Value{})
		// 存在错误直接返回
		if str := m[0].String(); str != "" {
			err = errors.New(str)
			return
		}
	}

	// 执行新增
	id, err = orm.NewOrm().Insert(object)

	// 新增成功执行新增之后的处理
	if err == nil {
		a := v.MethodByName("AfterInsert")
		if a.IsValid() {
			f.Call([]reflect.Value{})
		}
	}
	return
}

// 修改数据
func Update(object interface{}) (num int64, err error) {
	v := reflect.ValueOf(object)
	f := v.MethodByName("BeforeUpdate")
	// 执行修改之前
	if f.IsValid() {
		m := f.Call([]reflect.Value{})
		// 存在错误直接返回
		if str := m[0].String(); str != "" {
			err = errors.New(str)
			return
		}
	}

	// 执行修改
	num, err = orm.NewOrm().Update(object)

	// 执行修改之后
	if err == nil {
		a := v.MethodByName("AfterUpdate")
		if a.IsValid() {
			f.Call([]reflect.Value{})
		}
	}
	return
}

// 删除数据
func Delete(object interface{}) (num int64, err error) {
	// 获取反射信息
	v := reflect.ValueOf(object)
	f := v.MethodByName("BeforeDelete")
	// 执行删除之前
	if f.IsValid() {
		m := f.Call([]reflect.Value{})
		// 存在错误直接返回
		if str := m[0].String(); str != "" {
			err = errors.New(str)
			return
		}
	}

	// 执行删除
	num, err = orm.NewOrm().Delete(object)

	// 执行删除之后的处理
	if err == nil {
		a := v.MethodByName("AfterDelete")
		if a.IsValid() {
			f.Call([]reflect.Value{})
		}
	}

	return
}
