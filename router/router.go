package router

import (
	"net/http"

	"github.com/spf13/viper"
)

func Run() {
	port := viper.GetString("server.port")
	http.ListenAndServe(port, nil)
}
