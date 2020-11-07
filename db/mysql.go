package utils

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func Mysql(hostname string, port int, username string, password string, dbname string) (*sql.DB, error) {
	var err error
	Db, err = sql.Open("mysql",
		fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			username, password, hostname, port, dbname))
	if err != nil {
		panic(err.Error())
	}
	return Db, nil
}
