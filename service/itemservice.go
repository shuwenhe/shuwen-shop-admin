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

func GetItemCount() *common.Result {
	result := &common.Result{}
	totalRecord := dao.GetItemCount()
	if totalRecord >= 0 {
		result.Status = 200
		result.Data = totalRecord
		result.Msg = "Get the item count success!"
	} else {
		result.Status = 500
		result.Msg = "Get the item count fail!"
	}
	return result
}

func DeleteItemByID(id int) *common.Result {
	result := &common.Result{}
	if id <= 0 {
		result.Status = 400
		result.Msg = "ID is not null!"
	} else {
		dao.DeleteItemByID(id)
		result.Status = 200
		result.Msg = "Delete the item success!"

	}
	return result
}
