package main

import (
	"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/goincremental/negroni-sessions"
	"github.com/goincremental/negroni-sessions/cookiestore"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/unrolled/render"
	"log"
	"net/http"
)

var db, _ = gorm.Open("postgres", "user=postgres dbname=pqgotest password=1234 sslmode=disable")

type User struct {
	Id int64
	UserName string
	UserEmail string
	UserPassword string
}

func main() {

	defer db.Close()

	mux := http.NewServeMux()
	n := negroni.Classic()

	store := cookiestore.New([]byte("secret123"))
	n.Use(sessions.Sessions("my_session", store))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		SimplePage(w, r, "mainpage")
	})

	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			SimplePage(w, r, "login")
		} else if r.Method == "POST" {
			LoginPost(w, r)
		}
	})

	mux.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			SimplePage(w, r, "signup")
		} else if r.Method == "POST" {
			SignupPost(w, r)
		}
	})

	mux.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		Logout(w, r)
	})

	mux.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		SimpleAuthenticatedPage(w, r, "home")
	})

	mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		APIHandler(w, r)
	})

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	n.UseHandler(mux)
	n.Run(":3000")

}


func errHandler(err error) {
	if err != nil {
		log.Print(err)
	}
}


func SimplePage(w http.ResponseWriter, req *http.Request, template string) {
	r := render.New(render.Options{})
	r.HTML(w, http.StatusOK, template, nil)
}


func SimpleAuthenticatedPage(w http.ResponseWriter, req *http.Request, template string) {
	session := sessions.GetSession(req)
	sess := session.Get("useremail")

	if sess == nil {
		http.Redirect(w, req, "/notauthenticated", 301)
	}

	r := render.New(render.Options{})
	r.HTML(w, http.StatusOK, template, nil)
}

func LoginPost(w http.ResponseWriter, req *http.Request) {

	session := sessions.GetSession(req)

	username := req.FormValue("inputUsername")
	password := req.FormValue("inputPassword")

	var (
		email string
	)
	if username == password {
		log.Print("Good")
	}

	user := User{}
	db.Where(&User{UserEmail: username, UserPassword: password}).First(&user)

	// if err != nil {
	// 	log.Print(err)
	// 	http.Redirect(w, req, "/authfail", 301)
	// }

	session.Set("useremail", user.UserEmail)
	http.Redirect(w, req, "/home", 302)

}

func SignupPost(w http.ResponseWriter, req *http.Request) {

	username := req.FormValue("inputUsername")
	password := req.FormValue("inputPassword")
	email := req.FormValue("inputEmail")

	user := User{UserName: username, UserEmail: email, UserPassword: password}
	db.Create(&user)

	http.Redirect(w, req, "/login", 302)

}


func Logout(w http.ResponseWriter, req *http.Request) {

    session := sessions.GetSession(req)
    session.Delete("useremail") 
    http.Redirect(w, req, "/", 302)

}

func APIHandler(w http.ResponseWriter, req *http.Request) {

	data, _ := json.Marshal("{'API Test':'Works!'}")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(data)

}
