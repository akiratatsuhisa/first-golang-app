package db

import (
	"os"

	"github.com/akiratatsuhisa/first-golang-app/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Context *gorm.DB

func Connect() {
	var err error
	connectionStr, ok := os.LookupEnv("MYSQL_CONNECTION")

	if !ok {
		panic("missing MySQL connection string")
	}

	Context, err = gorm.Open(mysql.Open(connectionStr), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	Context.AutoMigrate(&models.User{}, &models.Role{}, &models.UserRole{}, &models.Todo{})
}
