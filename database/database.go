package database

import (
	models "learngo/app/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var connection *gorm.DB

func Connect() *gorm.DB {
	var err error
	connection, err = gorm.Open(mysql.Open(os.Getenv("DSN")), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // my current db does not allow this
	})
	if err != nil {
		panic("failed to connect database")
	}
	return connection
}

func Migrate() {
	connection.AutoMigrate(&models.Product{})
}

func DB() *gorm.DB {
	return connection
}
