package main

import (
	"NusaMeals_20/config"
)

func main() {
	// Inisialisasi koneksi database
	db := config.InitDB()

	// Migrasi tabel
	config.InitMigrate(db)
}
