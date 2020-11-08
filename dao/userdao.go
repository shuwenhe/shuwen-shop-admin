package dao

import (
	"log"

	"github.com/shuwenhe/shuwen-shop-admin/model"
	utils "github.com/shuwenhe/shuwen-shop/db"
)

func GetUser(name, password string) (*model.User, error) {
	log.Println("name-dao = *** = ", name)
	log.Println("password-dao = *** = ", password)
	sql := "select * from user where username=? and password=?"
	row := utils.Db.QueryRow(sql, name, password)
	user := &model.User{}
	log.Println("user = *** = ", user)
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Phone, &user.Email, &user.Created, &user.Updated)
	log.Println("user = *** = ", user)
	return user, nil
}
