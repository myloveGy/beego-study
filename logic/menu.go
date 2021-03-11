package logic

import (
	"time"

	"github.com/astaxie/beego/orm"

	"project/cache"
	"project/models"
)

type Menu struct {
	Id       int64   `json:"id"`
	Pid      int64   `json:"pid"`
	MenuName string  `json:"menu_name"`
	Icons    string  `json:"icons"`
	Url      string  `json:"url"`
	Children []*Menu `json:"child"`
}

func GetCacheMenu() []*Menu {
	// 第一步：缓存中获取
	menu, err := cache.Get("menu")
	if err != nil {
		if v, ok := menu.([]*Menu); ok {
			return v
		}
	}

	// 没有缓存数据信息 - 查询导航栏信息
	menuList := make([]*models.Menu, 0)
	orm.NewOrm().QueryTable(&models.Menu{}).Filter("status", 1).OrderBy("sort", "id").All(&menuList)
	data := FindParentMenu(0, menuList)
	if len(data) > 0 {
		cache.Put("menu", data, 43200*time.Second)
	}

	return data
}

func FindParentMenu(pid int64, list []*models.Menu) []*Menu {
	menuList := make([]*Menu, 0)
	for _, v := range list {
		if v.Pid == pid {
			menuList = append(menuList, &Menu{
				Id:       v.Id,
				Pid:      v.Pid,
				MenuName: v.MenuName,
				Url:      v.Url,
				Icons:    v.Icons,
				Children: make([]*Menu, 0),
			})
		}
	}

	for _, v := range menuList {
		children := FindParentMenu(v.Id, list)
		if len(children) > 0 {
			v.Children = children
		}
	}

	return menuList
}
