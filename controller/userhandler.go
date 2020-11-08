package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/shuwenhe/shuwen-shop-admin/service"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	result := service.GetUser(username, password)
	byteUser, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json:charset=utf-8")
	w.Write(byteUser)
}

func CheckUserName(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	log.Println("username-controller = *** = ", username)
	user, _ := service.CheckUserName(username)
	if user.ID > 0 {
		w.Write([]byte("用户名已存在！"))
	} else {
		w.Write([]byte("<font style = 'color:green'>用户名不存在！</font>"))
	}
	bsUser, _ := json.Marshal(user)
	w.Write(bsUser)
}
