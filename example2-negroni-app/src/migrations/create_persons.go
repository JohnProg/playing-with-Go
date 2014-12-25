package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"../models"
)

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var db, err = gorm.Open("postgres", "user=postgres dbname=pqgotest password=1234 sslmode=disable")
	PanicIf(err)
	db.CreateTable(data.Person{})
}