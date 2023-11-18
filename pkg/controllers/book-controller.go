package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/furkanadiiguzel/golang-bookstore/pkg/models"
	"github.com/furkanadiiguzel/golang-bookstore/pkg/utils"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r http.Request) {
	books := models.GetBook()
	res, _ := json.Marshal(books)
	fmt.Println("Endpoint Hit: GetBook")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	utils.ParseBody(r, &NewBook)
	createdBook := NewBook.CreateBooks()
	res, _ := json.Marshal(createdBook)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
