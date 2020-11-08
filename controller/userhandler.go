package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/shuwenhe/shuwen-shop-admin/service"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	password := r.FormValue("password")
	log.Println("name-controller = *** =", name)
	log.Println("password-controller = *** =", password)
	user, _ := service.GetUser(name, password)
	bsUser, _ := json.Marshal(user)
	w.Write(bsUser)
}
