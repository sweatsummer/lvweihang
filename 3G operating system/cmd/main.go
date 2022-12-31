package main

import (
	"3G_operating_system/api"
	"3G_operating_system/dao"
)

func main() {
	dao.InitDB()
	api.InitRouter()
}
