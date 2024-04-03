package configs

import (
	"prioritas1gorm/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	//Koneksi ke Database
	dsn := "root:@tcp(localhost:3306)/gorm1?charset=utf8mb4&parseTime=True&loc=Local"

	var err error

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	//Migrasi Skema Database
	AutoMigration()
}

func AutoMigration() {
	DB.AutoMigrate(&models.Package{}, &models.User{})
}
