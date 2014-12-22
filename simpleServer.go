package main

import(
  "net/http"
)

func main() {
  http.Handle("/", http.FileServer(http.Dir("./file.txt")))
  http.ListenAndServe(":8108", nil)
}
