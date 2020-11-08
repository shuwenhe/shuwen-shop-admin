package service

import (
	"github.com/shuwenhe/shuwen-shop-admin/common"
	"github.com/shuwenhe/shuwen-shop-admin/dao"
)

func GetItemByPage(rows, page int) (result common.Result) {
	items, _ := dao.GetItemByPage(rows, page)
	if len(items) > 0 {
		result.Status = 200
	} else {
		result.Status = 400
		result.Msg = "Get items fail!"
		return
	}
	result.Data = items
	result.Msg = "Get items success!"
	return
}
