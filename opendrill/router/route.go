package router

import (
	"../app/controllers/website"
	"../app/controllers/admin"
	"../app/controllers/client"
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
	
	//Website
	router.Handle("/register/", middleware.Handler(website.RegisterUser)).Methods("POST")

	//Administrator
	router.Handle("/admin/organizations/", middleware.Handler(admin.ListOrganizations)).Methods("GET")

	//Client

	//Organizator
	router.Handle("/client/organizator/{organizatorID}/organizations/", middleware.Handler(client.GetOrganizationsFromOrganizator)).Methods("GET")
	router.Handle("/client/organizator/{organizatorID}/add-organization/", middleware.Handler(client.AddOrganization)).Methods("POST")
	//Organization
	router.Handle("/client/organization/{organizationID}/", middleware.Handler(client.GetOrganization)).Methods("GET")
	router.Handle("/client/organization/{organizationID}/add-user/", middleware.Handler(client.AddUserToOrganzation)).Methods("POST")
	router.Handle("/client/organization/{organizationID}/delete-user/{userID}", middleware.Handler(client.DeleteUserToOrganzation)).Methods("DELETE")



	//Category
	//Update: db.organizations.update({name: "Mery2"}, {$set:{ruc:"123123123"}});
	//db.organizations.update( {_id: ObjectId("54c3e639b71b7f1fed000002")}, {$push:{listContacts: "prueba"}} )
	/*
	router.Handle("/categories/", middleware.Handler(controllers.ListCategories)).Methods("GET")
	router.Handle("/categories/", middleware.Handler(controllers.AddCategory)).Methods("POST")
	router.Handle("/categories/{categoryID}", middleware.Handler(controllers.GetCategory)).Methods("GET")
	router.Handle("/categories/{categoryID}", middleware.Handler(controllers.UpdateCategory)).Methods("PUT")
	router.Handle("/categories/{categoryID}", middleware.Handler(controllers.RemoveCategory)).Methods("DELETE")

	//Contact
	router.Handle("/contacts/", middleware.Handler(controllers.ListContacts)).Methods("GET")
	router.Handle("/contacts/", middleware.Handler(controllers.AddContact)).Methods("POST")
	router.Handle("/contacts/{contactID}", middleware.Handler(controllers.GetContact)).Methods("GET")
	router.Handle("/contacts/{contactID}", middleware.Handler(controllers.UpdateContact)).Methods("PUT")
	router.Handle("/contacts/{contactID}", middleware.Handler(controllers.RemoveContact)).Methods("DELETE")

	//List Contact
	router.Handle("/list-contacts/", middleware.Handler(controllers.ListListContacts)).Methods("GET")
	router.Handle("/list-contacts/", middleware.Handler(controllers.AddListContact)).Methods("POST")
	router.Handle("/list-contacts/{listContactID}", middleware.Handler(controllers.GetListContact)).Methods("GET")
	router.Handle("/list-contacts/{listContactID}", middleware.Handler(controllers.UpdateListContact)).Methods("PUT")
	router.Handle("/list-contacts/{listContactID}", middleware.Handler(controllers.RemoveListContact)).Methods("DELETE")

	//Designer
	router.Handle("/designers/", middleware.Handler(controllers.ListDesigners)).Methods("GET")
	router.Handle("/designers/", middleware.Handler(controllers.AddDesigner)).Methods("POST")
	router.Handle("/designers/{designerID}", middleware.Handler(controllers.GetDesigner)).Methods("GET")
	router.Handle("/designers/{designerID}", middleware.Handler(controllers.UpdateDesigner)).Methods("PUT")
	router.Handle("/designers/{designerID}", middleware.Handler(controllers.RemoveDesigner)).Methods("DELETE")

	//Template
	router.Handle("/designers/{designerID}/templates/", middleware.Handler(controllers.AddTemplateToDesigner)).Methods("POST")
	router.Handle("/designers/{designerID}/templates/", middleware.Handler(controllers.GetTemplatesFromDesigner)).Methods("GET")
	router.Handle("/designers/{designerID}/templates/{templateID}", middleware.Handler(controllers.GetTemplateFromDesigner)).Methods("GET")
	router.Handle("/designers/{designerID}/templates/{templateID}", middleware.Handler(controllers.UpdateTemplateFromDesigner)).Methods("PUT")
	router.Handle("/designers/{designerID}/templates/{templateID}", middleware.Handler(controllers.RemoveTemplateFromDesigner)).Methods("DELETE")

	//Organizator
	router.Handle("/organizator/", middleware.Handler(controllers.AllOrganizators)).Methods("GET")
	router.Handle("/organizator/", middleware.Handler(controllers.CreateOrganizator)).Methods("POST")
	router.Handle("/organizator/{organizatorID}", middleware.Handler(controllers.GetOrganizator)).Methods("GET")
	router.Handle("/organizator/{organizatorID}", middleware.Handler(controllers.UpdateOrganizator)).Methods("PUT")
	router.Handle("/organizator/{organizatorID}", middleware.Handler(controllers.RemoveOrganizator)).Methods("DELETE")

	//Organization
	router.Handle("/organizator/{organizatorID}/organizations/", middleware.Handler(controllers.AddOrganizationToOrganizator)).Methods("POST")
	router.Handle("/organizator/{organizatorID}/organizations/", middleware.Handler(controllers.GetOrganizationsFromOrganizator)).Methods("GET")
	router.Handle("/organizator/{organizatorID}/organizations/{organizationID}", middleware.Handler(controllers.GetOrganizationFromOrganizator)).Methods("GET")
	router.Handle("/organizator/{organizatorID}/organizations/{organizationID}", middleware.Handler(controllers.UpdateOrganizationFromOrganizator)).Methods("PUT")
	router.Handle("/organizator/{organizatorID}/organizations/{organizationID}", middleware.Handler(controllers.RemoveOrganizationFromOrganizator)).Methods("DELETE")
	*/
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileHandler))
	// router.PathPrefix("/").Handler(http.FileServer(http.Dir(config.Public)))
	http.Handle("/", router)
}
