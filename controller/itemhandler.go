package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/shuwenhe/shuwen-shop-admin/service"
)

func GetItemByPage(w http.ResponseWriter, r *http.Request) {
	rows, _ := strconv.Atoi(r.PostFormValue("rows"))
	page, _ := strconv.Atoi(r.PostFormValue("page"))
	pages := service.GetItemByPage(rows, page)
	byteResult, _ := json.Marshal(pages)
	w.Header().Set("Content-Type", "application/json:charset=utf-8")
	w.Write(byteResult)
}
