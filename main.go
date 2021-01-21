package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go-gin-mysql-boilerplate/Config"
	"go-gin-mysql-boilerplate/Models/Schema"
	"go-gin-mysql-boilerplate/Routes"
)
var err error
func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Schema.User{})
	r := Routes.SetupRouter()
	_ = r.Run()
}