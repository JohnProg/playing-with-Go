package controllers

import (
	models "../models"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	// "strconv"
)

// list of all of the books
// var books = make([]models.Book, 0)

func ListBooks(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	books, _ := models.AllBooks()
	if books == nil {
		return []models.Book{}, nil
	}
	return books, nil
}

func GetBook(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	// mux.Vars grabs variables from the path
	Id := mux.Vars(r)["id"]
	log.Println(Id)
	if len(Id) != 24 {
		return nil, &models.HandlerError{nil, "Id is not valid", http.StatusBadRequest}
	}

	err, b := models.GetBook(Id)

	if err != nil {
		return nil, &models.HandlerError{nil, "Could not find book " + Id, http.StatusNotFound}
	}

	return b, nil
}

func parseBookRequest(r *http.Request) (models.Book, *models.HandlerError) {
	// the book payload is in the request body
	data, e := ioutil.ReadAll(r.Body)

	if e != nil {
		return models.Book{}, &models.HandlerError{e, "Could not read request", http.StatusBadRequest}
	}
	// turn the request body (JSON) into a book object
	var payload models.Book

	e = json.Unmarshal(data, &payload)
	if e != nil {
		return models.Book{}, &models.HandlerError{e, "Could not parse JSON", http.StatusBadRequest}
	}

	return payload, nil
}

func AddBook(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
	payload, e := parseBookRequest(r)
	if e != nil {
		return nil, e
	}

	err := models.CreateBook(payload)
	log.Println(err)
	// we return the book we just made so the client can see the ID if they want
	return payload, nil
}

// func UpdateBook(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
// 	payload, e := parseBookRequest(r)
// 	if e != nil {
// 		return nil, e
// 	}

// 	_, index := getBookById(payload.Id)
// 	books[index] = payload
// 	return make(map[string]string), nil
// }

// func RemoveBook(w http.ResponseWriter, r *http.Request) (interface{}, *models.HandlerError) {
// 	param := mux.Vars(r)["id"]
// 	id, e := strconv.Atoi(param)
// 	if e != nil {
// 		return nil, &models.HandlerError{e, "Id should be an integer", http.StatusBadRequest}
// 	}
// 	// this is jsut to check to see if the book exists
// 	_, index := getBookById(id)

// 	if index < 0 {
// 		return nil, &models.HandlerError{nil, "Could not find entry " + param, http.StatusNotFound}
// 	}

// 	// remove a book from the list
// 	books = append(books[:index], books[index+1:]...)
// 	return make(map[string]string), nil
// }

// // searches the books for the book with `id` and returns the book and it's index, or -1 for 404
// func getBookById(id int) (models.Book, int) {
// 	for i, b := range books {
// 		if b.Id == id {
// 			return b, i
// 		}
// 	}
// 	return models.Book{}, -1
// }

// var id = 0

// // increments id and returns the value
// func getNextId() int {
// 	id += 1
// 	return id
// }
