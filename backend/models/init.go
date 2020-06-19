// Initialize models

package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var db *gorm.DB

// InitModels initializes database models
func InitModels() *gorm.DB {

	log.Print("Initializing Database.")
	err := godotenv.Load()

	if err != nil {
		log.Panic("env cannot be loaded.")
		log.Panic(err.Error())
	}

	if db == nil {
		username := os.Getenv("DB_USERNAME")
		password := os.Getenv("DB_PASSWORD")
		dbName := os.Getenv("DB_NAME")

		urlTemplate := "%s:%s@/%s?charset=utf8&parseTime=True&loc=Local"

		dbURL := fmt.Sprintf(urlTemplate, username, password, dbName)

		_db, err := gorm.Open("mysql", dbURL)
		if err != nil {
			log.Print("database initialization failed.")
			log.Panic(err.Error())
		}
		log.Print("Database initialization success.")

		db = _db
	}
	migrateModels(db)
	return db
}

// getAllModels are used for running auto migrations
func migrateModels(db *gorm.DB) {
	log.Print("Migrating models")
	db.AutoMigrate(&AccountType{})
	db.AutoMigrate(&ChartOfAccounts{})
	db.AutoMigrate(&Transactions{})
}
