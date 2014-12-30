package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request){
	//fmt.Fprintf(w, "Hello World")
	t, err := template.ParseFiles("templates/home.html","templates/header.html","templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "home", nil)
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
	name := r.FormValue("name")
	password := r.FormValue("password")
	re_password := r.FormValue("re_password")
}

func main(){

	fmt.Println("Listening on port :6969")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/save-register", saveUserHandler)
	http.ListenAndServe(":6969", nil)
}