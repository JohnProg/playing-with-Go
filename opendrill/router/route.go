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
	//Contact
	router.Handle("/contacts", middleware.Handler(controllers.ListContacts)).Methods("GET")
	router.Handle("/contacts", middleware.Handler(controllers.AddContact)).Methods("POST")
	router.Handle("/contacts/{id}", middleware.Handler(controllers.GetContact)).Methods("GET")
	router.Handle("/contacts/{id}", middleware.Handler(controllers.UpdateContact)).Methods("PUT")
	router.Handle("/contacts/{id}", middleware.Handler(controllers.RemoveContact)).Methods("DELETE")

	//Designer
	router.Handle("/designers", middleware.Handler(controllers.ListDesigners)).Methods("GET")
	router.Handle("/designers", middleware.Handler(controllers.AddDesigner)).Methods("POST")
	router.Handle("/designers/{designerID}", middleware.Handler(controllers.GetDesigner)).Methods("GET")
	router.Handle("/designers/{designerID}", middleware.Handler(controllers.UpdateDesigner)).Methods("PUT")
	router.Handle("/designers/{designerID}", middleware.Handler(controllers.RemoveDesigner)).Methods("DELETE")

	// //Template
	router.Handle("/designers/{designerID}/templates", middleware.Handler(controllers.AddTemplateToDesigner)).Methods("POST")
	router.Handle("/designers/{designerID}/templates", middleware.Handler(controllers.GetTemplatesFromDesigner)).Methods("GET")
	router.Handle("/designers/{designerID}/templates/{templateID}", middleware.Handler(controllers.GetTemplateFromDesigner)).Methods("GET")
	router.Handle("/designers/{designerID}/templates/{templateID}", middleware.Handler(controllers.UpdateTemplateFromDesigner)).Methods("PUT")
	router.Handle("/designers/{designerID}/templates/{templateID}", middleware.Handler(controllers.RemoveTemplateFromDesigner)).Methods("DELETE")

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileHandler))
	// router.PathPrefix("/").Handler(http.FileServer(http.Dir(config.Public)))
	http.Handle("/", router)
}
