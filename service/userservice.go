package service

import (
	"log"

	"github.com/shuwenhe/shuwen-shop-admin/dao"
	"github.com/shuwenhe/shuwen-shop-admin/model"
)

func GetUser(username, password string) (*model.User, error) {
	log.Println("name-service = *** =", username)
	log.Println("password-service = *** =", password)
	user, err := dao.GetUser(username, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func CheckUserName(username string) (*model.User, error) {
	log.Println("username-service = *** = ", username)
	user, err := dao.CheckUserName(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}
