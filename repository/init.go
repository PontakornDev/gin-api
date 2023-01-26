package repository

import (
	"fmt"
	"os"
	"pontakorn322/demo_gin_api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
	// config := utils.ReadConfigs()
	db, err := gorm.Open(mysql.Open(os.Getenv("MYSQLDB")), &gorm.Config{})
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
