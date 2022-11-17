package controllers

import (
	"github.com/nanichang/bookstore/pkg/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/nanichang/bookstore/pkg/utils"
)

var NewBook models.Book

func GetBook(res http.ResponseWriter, req *http.Request) {
	newBooks := models.GetAllBooks()
	response, _ := json.Marshal(newBooks)
	res.Header().Set("Content-Type", "pkglication/json")
	res.WriteHeader(http.StatusOK)
	res.Write(response)
}

func GetBookById(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId,0,0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	bookDetails, _ := models.GetBookById(ID)
	response, _ := json.Marshal(bookDetails)
	res.Header().Set("Content-Type", "pkglication/json")
	res.WriteHeader(http.StatusOK)
	res.Write(response)
}

func CreateBook(res http.ResponseWriter, req *http.Request) {
	bookModel := &models.Book{}
	utils.ParseBody(req, bookModel)
	book := bookModel.CreateBook()
	response, _ := json.Marshal(book)
	res.WriteHeader(http.StatusOK)
	res.Write(response)
}

func DeleteBook(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId,0,0)
	if err != nil {
		fmt.Println("error while parsing id")
	}
	book := models.DeleteBook(ID)
	response, _ := json.Marshal(book)
	res.Header().Set("Content-Type", "pkglication/json")
	res.WriteHeader(http.StatusOK)
	res.Write(response)
}

func UpdateBook(res http.ResponseWriter, req *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(req, updateBook)
	vars := mux.Vars(req)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId,0,0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	bookDetails, db := models.GetBookById(ID)
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

	response, _ := json.Marshal(bookDetails)
	res.Header().Set("Content-Type", "pkglication")
	res.WriteHeader(http.StatusOK)
	res.Write(response)
}