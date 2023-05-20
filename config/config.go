package config

import (
	"fmt"

	"NusaMeals_20/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	config := map[string]string{
		"DB_Username": "root",
		"DB_Password": "yuni260200",
		"DB_Port":     "3306",
		"DB_Host":     "localhost",
		"DB_Name":     "NusaMeals_20",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config["DB_Username"],
		config["DB_Password"],
		config["DB_Host"],
		config["DB_Port"],
		config["DB_Name"])

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	InitMigrate(DB) // Memanggil InitMigrate dengan parameter DB
	return DB
}

func InitMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
		&models.Level{},
		&models.Menu{},
		&models.Category{},
		&models.Food{},
		&models.Drink{},
		&models.Order{},
		&models.OrderCancel{},
		&models.Reservation{},
		&models.ListTable{},
		&models.Report{},
		&models.Customer{},
		&models.CustomerEwalletAccount{},
		&models.CustomerCardAccount{},
		&models.Payment{},
		&models.PaymentType{},
		&models.CardType{},
		&models.EwalletType{},
	)
	if err != nil {
		panic(err)
	}
}
