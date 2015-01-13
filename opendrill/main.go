package main

import (
	"./app"
	"log"
	"net/http"
)

func main() {
	var configPath = "config.json"

	a, err := app.NewApp(configPath)
	if err != nil {
		panic(err)
	}

	defer func() {
		a.Connection.Session.Close()
	}()

	log.Printf("Running on port ", a.Config.Port)
	if err := http.ListenAndServe(a.Config.Port, nil); err != nil {
		log.Println(err.Error())
	}
}
