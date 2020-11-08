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
	log.Println("name-controller = *** =", username)
	log.Println("password-controller = *** =", password)
	user, _ := service.GetUser(username, password)
	bsUser, _ := json.Marshal(user)
	w.Write(bsUser)
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
