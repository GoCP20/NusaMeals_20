package main

import (
	"mini-project/config"
	route "mini-project/routes"
	"mini-project/util"

	"github.com/go-playground/validator/v10"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	config.InitDB()
	config.InitialMigration()
}

func main() {

	e := route.New()
	e.Validator = &util.CustomValidator{
		Validator: validator.New(),
	}
	e.Logger.Fatal(e.Start(":8080"))
}
