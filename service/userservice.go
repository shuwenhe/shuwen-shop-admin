package service

import (
	"log"

	"github.com/shuwenhe/shuwen-shop-admin/dao"
	"github.com/shuwenhe/shuwen-shop-admin/model"
)

func GetUser(name, password string) (*model.User, error) {
	log.Println("name-service = *** =", name)
	log.Println("password-service = *** =", password)
	user, err := dao.GetUser(name, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
