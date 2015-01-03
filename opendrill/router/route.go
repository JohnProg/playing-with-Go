package router

import (
	"../app/controllers"
	"../app/middlewares"
	"github.com/gorilla/mux"
	"net/http"
)

func Init() {
	// handle all requests by serving a file of the same name
	fileHandler := http.FileServer(http.Dir("../public/"))

	// setup routes
	router := mux.NewRouter()
	router.Handle("/", http.RedirectHandler("/static/", 302))
	router.Handle("/books", middlewares.Handler(controllers.ListBooks)).Methods("GET")
	router.Handle("/books", middlewares.Handler(controllers.AddBook)).Methods("POST")
	router.Handle("/books/{id}", middlewares.Handler(controllers.GetBook)).Methods("GET")
	router.Handle("/books/{id}", middlewares.Handler(controllers.UpdateBook)).Methods("PUT")
	router.Handle("/books/{id}", middlewares.Handler(controllers.RemoveBook)).Methods("DELETE")

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileHandler))
	// router.PathPrefix("/").Handler(http.FileServer(http.Dir(config.Public)))
	http.Handle("/", router)
}
