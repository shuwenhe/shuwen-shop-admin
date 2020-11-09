package dao

import (
	"fmt"
	"testing"
)

func testGetCategoryByID(t *testing.T) {
	category, _ := GetCategoryByID(1)
	fmt.Println("category = *** = ", category)
}
