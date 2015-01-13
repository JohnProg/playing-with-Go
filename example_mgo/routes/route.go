package router

import (
	"../handlers"
	"../middlewares"
	"github.com/gorilla/mux"
	"net/http"
)

func Init() {
	// handle all requests by serving a file of the same name
	fileHandler := http.FileServer(http.Dir("client/"))

	// setup routes
	router := mux.NewRouter()
	router.Handle("/", http.RedirectHandler("/static/", 302))
	router.Handle("/books", middleware.Handler(controllers.ListBooks)).Methods("GET")
	router.Handle("/books", middleware.Handler(controllers.AddBook)).Methods("POST")
	router.Handle("/books/{id}", middleware.Handler(controllers.GetBook)).Methods("GET")
	router.Handle("/books/{id}", middleware.Handler(controllers.UpdateBook)).Methods("PUT")
	router.Handle("/books/{id}", middleware.Handler(controllers.RemoveBook)).Methods("DELETE")

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileHandler))
	http.Handle("/", router)
}
