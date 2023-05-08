package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

// ConnectDB connects to the database
func ConnectDB() {
	database, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("Failed to connect to database!")
	}
	database.AutoMigrate(&Lead{})
	db = database
}

func CloseDB() {
	defer db.Close()
}

func GetDB() *gorm.DB {
	return db
}