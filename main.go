package main

import (
	"github.com/Tasrifin/bookingtogo-golang/config"
)

func main() {
	config.InitEnv()
	config.InitRoute()
}
