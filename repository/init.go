package repository

import (
	"fmt"
	"pontakorn322/demo_gin_api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
	dsn := "root:@tcp(localhost:3306)/test2?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
