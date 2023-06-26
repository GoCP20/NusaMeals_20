package main

import (
	"reglog/internal/common/config"
	"reglog/seeder/data"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type Seeder struct {
	Seeder interface{}
}

func RegisterSeeder(db *gorm.DB) []Seeder {
	return []Seeder{
		{Seeder: data.UserFaker(db)},
	}
}

func main() {
	config.InitDB()
	db := config.DB
	defer db.Close()

	for _, seeder := range RegisterSeeder(db) {
		err := db.Debug().Create(seeder.Seeder).Error
		if err != nil {
			logrus.Fatal("Cannot run seeder: ", err.Error())
		} else {
			logrus.Info("Success run seeder")
		}
	}
}
