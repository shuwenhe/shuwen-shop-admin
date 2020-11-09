package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/shuwenhe/shuwen-shop-admin/service"
)

func GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PostFormValue("id"))
	result := service.GetCategoryByID(id)
	byteResult, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json:charset=utf-8")
	w.Write(byteResult)
}
