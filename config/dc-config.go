package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabaseConnection() *gorm.DB{
	err := godotenv.Load()
	if err != nil{
		panic("failed to load environment variables")
	}

	db_user := os.Getenv("db_user")
	db_password := os.Getenv("db_pass")
	dbHost := os.Getenv("db_host")
	db_name := os.Getenv("db_name")


	dsn := fmt.Sprint("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", db_user, db_password, dbHost, db_name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to create a connection to database")
	}

	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	
}