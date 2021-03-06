package main

import (
  "github.com/codegangsta/negroni"
  "github.com/goincremental/negroni-sessions"
  "github.com/goincremental/negroni-sessions/cookiestore"
  "net/http"
)

func main() {
  n := negroni.Classic()

  store := cookiestore.New([]byte("secret123"))  
  n.Use(sessions.Sessions("my_session", store))

  mux := http.NewServeMux()
  mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
    session := sessions.GetSession(req)
    session.Set("hello", "world")
  })

  n.UseHandler(mux)
  n.Run(":3000")
}