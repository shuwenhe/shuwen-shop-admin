package dao

import (
	"log"
	"testing"
)

func TestGetUser(t *testing.T) {
	user, _ := GetUser("shuwen", "123456")
	log.Println("user = *** =", user)
}
