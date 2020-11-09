package dao

import (
	"fmt"
	"log"
	"testing"
)

func testGetItemByPage(t *testing.T) {
	items, _ := GetItemByPage(4, 2)
	for _, item := range items {
		fmt.Println("item = *** =", item)
	}
}

func TestGetItemCount(t *testing.T) {
	totalRecord := GetItemCount()
	log.Println("totalRecord = *** = ", totalRecord)
}
