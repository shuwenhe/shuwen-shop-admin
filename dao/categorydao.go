package dao

import (
	utils "github.com/shuwenhe/shuwen-shop-admin/db"
	"github.com/shuwenhe/shuwen-shop-admin/model"
)

func GetCategoryByID(id int) (*model.Category, error) {
	category := &model.Category{}
	sql := "select * from admin_item_cat where id=?"
	row := utils.Db.QueryRow(sql, id)
	row.Scan(&category.ID, &category.ParentID, &category.Name, &category.Status, &category.SortOrder, &category.IsParent, &category.Created, &category.Updated)
	return category, nil
}
