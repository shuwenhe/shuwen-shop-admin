package dao

import (
	"log"

	utils "github.com/shuwenhe/shuwen-shop-admin/db"
	"github.com/shuwenhe/shuwen-shop-admin/model"
)

func GetUser(username, password string) (*model.User, error) {
	log.Println("name-dao = *** = ", username)
	log.Println("password-dao = *** = ", password)
	sql := "select * from users where username=? and password=?"
	log.Println("sql = *** = ", sql)
	row := utils.Db.QueryRow(sql, username, password)
	log.Println("row = *** = ", row)
	user := &model.User{}
	log.Println("user = *** = ", user)
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	log.Println("user = *** = ", user)
	return user, nil
}

func GetUser2(username string, password string) (*model.User, error) {
	sql := "select id,username,password,email from users where username = ? and password = ?"
	row := utils.Db.QueryRow(sql, username, password)
	user := &model.User{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
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
