package dao

import (
	utils "github.com/shuwenhe/shuwen-shop-admin/db"
	"github.com/shuwenhe/shuwen-shop-admin/model"
)

func GetItemByPage(rows, page int) ([]*model.Book, error) {
	sql := "select * from books limit ?,?"
	itemRows, err := utils.Db.Query(sql, rows*(page-1), rows)
	if err != nil {
		return nil, err
	}
	var items []*model.Book
	for itemRows.Next() {
		item := &model.Book{}
		itemRows.Scan(&item.ID, &item.Title, &item.Author, &item.Price, &item.Sales, &item.Stock, &item.ImgPath)
		items = append(items, item)
	}
	return items, nil
}
