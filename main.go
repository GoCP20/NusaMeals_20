package main

import "github.com/go-playground/validator"

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
