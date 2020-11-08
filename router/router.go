package router

import (
	"net/http"

	"github.com/shuwenhe/shuwen-shop-admin/controller"
	"github.com/spf13/viper"
)

func Run() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))

	http.HandleFunc("/login", controller.GetUser)
	http.HandleFunc("/checkUserName", controller.CheckUserName)
	http.HandleFunc("/getItemByPage", controller.GetItemByPage)

	port := viper.GetString("server.port")
	http.ListenAndServe(port, nil)
}
