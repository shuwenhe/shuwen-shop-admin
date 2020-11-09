package service

import (
	"github.com/shuwenhe/shuwen-shop-admin/common"
	"github.com/shuwenhe/shuwen-shop-admin/dao"
)

func GetCategoryByID(id int) *common.Result {
	result := &common.Result{}
	category, _ := dao.GetCategoryByID(id)
	if category.ID > 0 {
		result.Status = 200
		result.Data = category
		result.Msg = "Get the category success!"
	} else {
		result.Status = 500
		result.Msg = "Get the category fail!"
	}
	return result
}
