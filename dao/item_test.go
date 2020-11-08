package dao

import (
	"fmt"
	"testing"
)

func testGetItemByPage(t *testing.T) {
	items, _ := GetItemByPage(4, 2)
	for _, item := range items {
		fmt.Println("item = *** =", item)
	}
}
