package controllers

import (
	"encoding/json"
	"github.com/Golang-Personal-Projects/Go-Projects/03-Go-Mysql-API/pkg/models"
	"github.com/Golang-Personal-Projects/Go-Projects/03-Go-Mysql-API/pkg/utils"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

var NewBook models.Book

// GetBooks Get all books
func GetBooks(w http.ResponseWriter, req *http.Request) {

	newBook := models.GetAllBooks()

	res, err := json.Marshal(newBook)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// GetBookByID  get book by their ids
func GetBookByID(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		log.Fatal(err)
	}
	bookDetails, _ := models.GetBookByID(ID)

	res, err := json.Marshal(bookDetails)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// CreateBook  Create a new book
func CreateBook(w http.ResponseWriter, req *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(req, CreateBook)

	b := CreateBook.CreateBook()
	res, err := json.Marshal(b)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// DeleteBook Delete a book by Id
func DeleteBook(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		log.Fatal(err)
	}
	book := models.DeleteBook(ID)
	res, err := json.Marshal(book)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// UpdateBook - Update a book by Id
func UpdateBook(w http.ResponseWriter, req *http.Request) {

	var updateBook models.Book
	utils.ParseBody(req, updateBook)
	vars := mux.Vars(req)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		log.Fatal(err)
	}

	bookDetails, db := models.GetBookByID(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(&bookDetails)
	res, err := json.Marshal(bookDetails)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
