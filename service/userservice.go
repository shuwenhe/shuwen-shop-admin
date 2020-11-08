package service

import (
	"log"

	"github.com/shuwenhe/shuwen-shop-admin/common"
	"github.com/shuwenhe/shuwen-shop-admin/dao"
	"github.com/shuwenhe/shuwen-shop-admin/model"
)

func GetUser(username, password string) (result common.Result) {
	user, _ := dao.GetUser(username, password)
	if user.ID > 0 {
		result.Status = 200
	} else {
		result.Status = 400
	}
	result.Data = user
	result.Msg = "Login success!"
	return
}

func CheckUserName(username string) (*model.User, error) {
	log.Println("username-service = *** = ", username)
	user, err := dao.CheckUserName(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}
