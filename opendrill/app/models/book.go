package models

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"log"
)

// error response contains everything we need to use http.Error
type HandlerError struct {
	Error   error
	Message string
	Code    int
}

// book model
type Book struct {
	Id     bson.ObjectId `bson:"_id"`
	Title  string        `json:"title"`
	Author string        `json:"author"`
}

func AllBooks() (msgs []Book) {
	_ = books.
		Find(bson.M{}).
		All(&msgs)
	log.Println(msgs)
	return
}

// Save inserts a new book into MongoDB
func CreateBook(book Book) error {

	book.Id = bson.NewObjectId()

	if err := books.Insert(book); err != nil {
		return fmt.Errorf("Error creating new book: %v", err)
	}
	log.Println(book)
	return nil
}
