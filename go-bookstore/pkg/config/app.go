package config
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// ConnectDB connects to the database
func ConnectDB() {
	database, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/go_bookstore_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Failed to connect to database!")
	}
	database.AutoMigrate(&Book{})
	db = database
}

func CloseDB() {
	defer db.Close()
}

func GetDB() *gorm.DB {
	return db
}



