package router

import (
	"net/http"

	"github.com/shuwenhe/shuwen-shop-admin/controller"
	"github.com/spf13/viper"
)

func Run() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))
	RouterGroup()
	port := viper.GetString("server.port")
	http.ListenAndServe(port, nil)
}

func RouterGroup() {
	UserGroup()
	ItemGroup()
	CategoryGroup()
}

func UserGroup() {
	http.HandleFunc("/login", controller.GetUser)
	http.HandleFunc("/checkUserName", controller.CheckUserName)
}

func ItemGroup() {
	http.HandleFunc("/getItemByPage", controller.GetItemByPage)
	http.HandleFunc("/getItemCount", controller.GetItemCount)
	http.HandleFunc("/deleteItemByID", controller.DeleteItemByID)
	http.HandleFunc("/instockItemByID", controller.InstockItemByID)
}

func CategoryGroup() {
	http.HandleFunc("/getCategoryByID", controller.GetCategoryByID)
}
