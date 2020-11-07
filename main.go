package main

import (
	"os"

	"github.com/shuwenhe/shuwen-shop-admin/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		println("Start fail:", err.Error())
		os.Exit(-1)
	}
}
