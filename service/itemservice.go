package service

import (
	"github.com/shuwenhe/shuwen-shop-admin/common"
	"github.com/shuwenhe/shuwen-shop-admin/dao"
)

func GetItemByPage(rows, page int) *common.Page {
	pages := &common.Page{}
	items, _ := dao.GetItemByPage(rows, page)
	if items != nil {
		pages.Rows = items
		pages.Total = len(items)
	}
	return pages
}
