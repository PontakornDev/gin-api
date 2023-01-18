package repository

import (
	"fmt"
	"pontakorn322/demo_gin_api/models"
	"pontakorn322/demo_gin_api/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
	config := utils.ReadConfigs()
	db, err := gorm.Open(mysql.Open(config.DatabaseConnection), &gorm.Config{})
	if err != nil {
		return err
	}
	fmt.Println("Database is Connected")
	DB = db
	return err
}

func MigrationDB() {
	DB.AutoMigrate(&models.Product{})
}
