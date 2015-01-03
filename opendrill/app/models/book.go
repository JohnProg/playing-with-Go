package models

import (
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
	Id     bson.ObjectId `bson:"_id" json:"id"`
	Title  string        `json:"title"`
	Author string        `json:"author"`
}

func AllBooks() (books2 []Book, err error) {
	err = books.
		Find(nil).
		All(&books2)
	return
}

// Save inserts a new book into MongoDB
func CreateBook(book Book) error {
	// First, let's get a new id
	obj_id := bson.NewObjectId()
	book.Id = obj_id

	if err := books.Insert(book); err != nil {
		return err
	}
	return nil
}

func GetBook(Id string) (err error, book Book) {
	bid := bson.ObjectIdHex(Id)
	err = books.
		FindId(bid).
		One(&book)
	log.Println(book)
	return
}
