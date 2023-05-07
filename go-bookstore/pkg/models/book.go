package models

import (
	"github.com/jinzhu/gorm"
	"github.com/Jonathansoufer/go-bookstore/pkg/config"
)

var db *gorm.DB

// Book struct
type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

// CreateBook creates a book
func (book *Book) CreateBook() *Book{
	db.NewRecord(book)
	db.Create(&book)
	return book
}

// GetBook by id
func GetBookById(id int64) (*Book, *gorm.DB){
	var getBook Book
	db := db.Where("ID = ?", id).Find(&getBook)
	return &getBook, db
}

// GetAllBooks gets all books
func GetAllBooks() []Book{
	var books []Book
	db := db.Find(&books)
	return books
}

// UpdateBook updates a book
func (book *Book) UpdateBook() *Book{
	db.Save(&book)
	return book
}

// DeleteBook deletes a book
func DeleteBook(id int64) Book{
	var book Book
	db.Where("ID = ?", id).Delete(book)
	return book
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}