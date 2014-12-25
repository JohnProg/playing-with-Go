package main

import (
	"./models"
	"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/goincremental/negroni-sessions"
	"github.com/goincremental/negroni-sessions/cookiestore"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/unrolled/render"
	"io/ioutil"
	"log"
	"net/http"
)

var db gorm.DB
var r = render.New(render.Options{
        IndentJSON: true,
    })
// func SetupDB(db gorm.DB) gorm.DB {
// 	var err error
// 	var config data.Config
// 	var configraw, _ = ioutil.ReadFile("config.json")	
// 		json.Unmarshal(configraw, &config)
// 	db, err = gorm.Open(config.Database.Server, "user=" + config.Database.User + " dbname=" + config.Database.Name +" password="+ config.Database.Password +" sslmode=disable")
// 	PanicIf(err)
// 	return db
// }

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	mux := http.NewServeMux()
	n := negroni.Classic()

	store := cookiestore.New([]byte("secret123"))
	n.Use(sessions.Sessions("my_session", store))
	// n.Map(SetupDB())
	
	var config data.Config
	var configraw, _ = ioutil.ReadFile("config.json")	
		json.Unmarshal(configraw, &config)
	db, _ = gorm.Open(config.Database.Server, "user=" + config.Database.User + " dbname=" + config.Database.Name +" password="+ config.Database.Password +" sslmode=disable")

	// db := SetupDB(db)
	defer db.Close()

	mux.HandleFunc("/", Index)
	mux.HandleFunc("/api1/", func(w http.ResponseWriter, req *http.Request) {
        r.JSON(w, http.StatusOK, map[string]string{"welcome": "This is rendered JSON!"})
    })

	 mux.HandleFunc("/profile/", func(w http.ResponseWriter, req *http.Request) {
        w.Write([]byte("Welcome, visit sub pages now."))
    })	
	mux.HandleFunc("/data/", func(w http.ResponseWriter, req *http.Request) {
        r.Data(w, http.StatusOK, []byte("Some binary data here."))
    })

    mux.HandleFunc("/json/", func(w http.ResponseWriter, req *http.Request) {
        r.JSON(w, http.StatusOK, map[string]string{"hello": "json"})
    })

    mux.HandleFunc("/jsonp", func(w http.ResponseWriter, req *http.Request) {
        r.JSONP(w, http.StatusOK, "callbackName", map[string]string{"hello": "jsonp"})
    })

    mux.HandleFunc("/xml/", func(w http.ResponseWriter, req *http.Request) {
        r.XML(w, http.StatusOK, data.ExampleXml{One: "hello", Two: "xml"})
    })
    mux.HandleFunc("/html/", func(w http.ResponseWriter, req *http.Request) {
        // Assumes you have a template in ./templates called "example.tmpl"
        // $ mkdir -p templates && echo "<h1>Hello {{.}}.</h1>" > templates/example.tmpl
        r.HTML(w, http.StatusOK, "example", nil)
    })

	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			SimpleRedirectPage(w, r, "login")
		} else if r.Method == "POST" {
			LoginPost(w, r)
		}
	})

	mux.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			SimpleRedirectPage(w, r, "signup")
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

func Index(w http.ResponseWriter, req *http.Request) {
	users := []data.User{}
	db.Find(&users)
	r.HTML(w, http.StatusOK, "mainpage", users)
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

func SimpleRedirectPage(w http.ResponseWriter, req *http.Request, template string) {
	session := sessions.GetSession(req)
	sess := session.Get("useremail")
	log.Println(sess)
	if sess == nil {
		SimplePage(w, req, template)
	} else {
		http.Redirect(w, req, "/home", 302)
	}
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

	user := data.User{}
	db.Where(&data.User{UserEmail: username, UserPassword: password}).First(&user)

	if user.UserPassword != password {
		http.Redirect(w, req, "/authfail", 301)
	}

	session.Set("useremail", user.UserEmail)
	http.Redirect(w, req, "/home", 302)

}

func SignupPost(w http.ResponseWriter, req *http.Request) {

	username := req.FormValue("inputUsername")
	password := req.FormValue("inputPassword")
	email := req.FormValue("inputEmail")

	user := data.User{UserName: username, UserEmail: email, UserPassword: password}
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