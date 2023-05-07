package controllers

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/Jonathansoufer/go-bookstore/pkg/models"
	"github.com/Jonathansoufer/go-bookstore/pkg/utils"
	"strconv"
	"encoding/json"
	"fmt"
)

var NewBook models.Book

// CreateBook creates a book
func CreateBook(w http.ResponseWriter, r *http.Request){
	utils.ParseBody(r, &NewBook)
	book := NewBook.CreateBook()
	res, _ := json.Marshal(book)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// GetBook gets a book
func GetBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// GetAllBooks gets all books
func GetAllBooks(w http.ResponseWriter, r *http.Request){
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// UpdateBook updates a book
func UpdateBook(w http.ResponseWriter, r *http.Request){
	utils.ParseBody(r, &NewBook)
	book := NewBook.UpdateBook()
	res, _ := json.Marshal(book)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// DeleteBook deletes a book
func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}