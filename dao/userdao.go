package dao

import (
	"log"

	utils "github.com/shuwenhe/shuwen-shop-admin/db"
	"github.com/shuwenhe/shuwen-shop-admin/model"
)

func GetUser(username, password string) (*model.User, error) {
	sql := "select * from admin_user where username=? and password=? or phone=? and password=? or email=? and password=?"
	row := utils.Db.QueryRow(sql, username, password, username, password, username, password)
	user := &model.User{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Phone, &user.Email, &user.Created, &user.Updated)
	return user, nil
}

func CheckUserName(username string) (*model.User, error) {
	log.Println("username-dao = *** = ", username)
	sql := "select id,username,password,email from users where username = ?"
	row := utils.Db.QueryRow(sql, username)
	log.Println("row = *** = ", row)
	user := &model.User{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	return user, nil
}
