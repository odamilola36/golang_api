package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/odamilola36/golang_api/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// openConnection to the database from the app
func SetupDatabaseConnection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic("failed to load environment variables")
	}

	db_user := os.Getenv("db_user")
	db_password := os.Getenv("db_pass")
	dbHost := os.Getenv("db_host")
	db_name := os.Getenv("db_name")


	dsn := fmt.Sprint("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", db_user, db_password, dbHost, db_name)

	//TODO format the dsn string
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to create a connection to database")
	}

	db.AutoMigrate(entity.User{}, &entity.Book{})
	return db
}

//close connection to the database form your app
func CloseDatabaseConnection(db *gorm.DB) {
	dbSql, err := db.DB()

	if err != nil {
		panic("Failed to create a connection from database")
	}

	dbSql.Close()
}
