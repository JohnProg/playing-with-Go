package main

import (
  "github.com/codegangsta/negroni"
  "net/http"
  "fmt"
)

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "Hello world john!")
  })
  n := negroni.Classic()
  n.UseHandler(mux)
  n.Run(":3000")
}