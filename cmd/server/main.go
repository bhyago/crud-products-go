package main

import "github.com/bhyago/crud-products-go/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBHost)
}
