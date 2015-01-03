package main

import (
	"fmt"
	"html/template"
	"net/http"
	"./models"
	"./utils"
)

var users map[string]*models.User

func homeHandler(w http.ResponseWriter, r *http.Request){
	//fmt.Fprintf(w, "Hello World")
	t, err := template.ParseFiles("templates/home.html","templates/header.html","templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Println(users)

	t.ExecuteTemplate(w, "home", users)
}


func loginHandler(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("templates/login.html", "templates/header.html", "templates/footer.html")
	if err != nil{
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "login", nil)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/register.html", "templates/header.html", "templates/footer.html")
	if err !=  nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "register", nil)
}

func saveUserHandler(w http.ResponseWriter, r *http.Request){
	id := r.FormValue("id")
	email := r.FormValue("email")
	password := r.FormValue("password")
	re_password := r.FormValue("re_password")

	var user *models.User
	if id != ""{
		user = users[id]
		if password == re_password {
			user.Email = email
			user.Password = password
		}
	} else {
		id := generator.GenerateId()

		if password == re_password {
			user := models.NewUser(id, email, password)
			users[user.Id] = user
		}
	}

	http.Redirect(w, r ,"/", 302)
}

func editHandler(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("templates/edit.html", "templates/header.html", "templates/footer.html" )

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	id := r.FormValue("id")
	user, found := users[id]
	if !found{
		http.NotFound(w, r)
	}

	fmt.Println("encontrado", user)

	t.ExecuteTemplate(w, "edit", user)
}

func main(){

	fmt.Println("Listening on port :6969")

	users = make(map[string]*models.User, 0)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/save-register", saveUserHandler)
	http.HandleFunc("/edit", editHandler)
	http.ListenAndServe(":6969", nil)
}